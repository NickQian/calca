% 时域左边加窗，固定分析比如750点。
% 1) 以适应新上市股票，eg. 兆易创新只有800个交易日
% 2) 以用于不同指数在时域对比(时域数据同步)

%本程序只处理 1)

%pkg load signal

%======================  参数  ===================================
function [k_org, k_prc] = fsf_zph_ltwin(fn, fc_in, len_in)   %'kline.txt'


Nsig_FixLen = len_in  %len_in                      % 分析3年       
fc = fc_in %fc_in                         % 50/60/70    X:100   lowpass cutoff frequency in Hz
Len_W_FD_R_Ratio = 1/7             % 频域主窗右侧加窗长度,为整个box的1/Len_W_FD_R_DIV
Len_W_TD_L_Ratio = 1/10            % 多长比例的数据加窗


%FFT processing parameters:
L =  2047;  M=L         % FIR filter length in taps  % nominal window length 
Nfft = 2^(ceil(log2(M+L-1)))    % FFT Length. L=31-> Nfft=64

fs = 4000                       % sampling rate in Hz 
T  = 1;   dt = 1.0/fs;  
Nfft_fd = 16384   %8192    % 4096

%t = linspace(0, T,    fs);         % 生成2048个点。这些点的间距为 (x2-x1)/(n-1)。


%---- 时域窗 幅值修正系数-----
C_wc = 0.5396      % hanning-0.4995  hamming-0.5396   blackman-0.4196
C_wc_mul = 1/C_wc

%---- 频域窗 幅值修正系数-----
C_wc_FD = 1.313
C_wc_FD_mul = 1/C_wc_FD;

%========================== 输入信号 ==============================
%----- 导入数据 ------
d  = importdata(fn);    %(fn)   ("kline_sh.txt") % fscanf 和 textscan       %[x1, x2, x3, x4] = textread('kline.txt', '%2s  %.3f  %.3f  %.3f')
kline_all  = d';         %d.data  d.textdata  d.rowheaders
k1_lastDay = kline_all(length(kline_all))   
Nsig_max   = length(kline_all)                


%------------- 截取右侧Nsig_FixLen长度信号 -----------
kline = kline_all(Nsig_max-Nsig_FixLen+1 : Nsig_max)
k_org = kline;                   % @@@ %% for function output
L_Kline = length(kline)

%----------------- DC Remove ------------------
dc_sum = 0;
for k = 1 : L_Kline
  dc_sum = dc_sum + kline(k);    
endfor
dc = dc_sum/L_Kline

kline_noDc = [];
for k = 1:L_Kline
  kline_noDc(k) = kline(k) - dc;
endfor


%----------------- TG(tree growth) Removal -------------------
% 选定两个bot，找出斜率
bot1 = kline_noDc(1)
bot2 = kline_noDc(L_Kline)
TG_high = (bot2 + bot1)/2;

% 画出TG
% tg = (-(TG_high -1)/2 : (TG_high -1)/2)
tg_pre = linspace(bot1,  bot2,   L_Kline   );
tg     = tg_pre - TG_high;

% TG remove
%kline_NoDcTg = kline_noDc - tg   % 移除TG
kline_NoDcTg = kline_noDc;        % 不移除TG

%--------------------- 右侧移零 ------------------------------
kline_NoDcTg_r0 = kline_NoDcTg - kline_NoDcTg(length(kline_NoDcTg))


%-------------------- 预处理结果绘图 ---------------------------
% 绘图读入数据
figure(1)
day = 1:1:L_Kline;

subplot(311);
plot(day, kline);                   
xlabel('time/day');ylabel('Amplitude');title('kline');

subplot(312);
plot(day, kline_noDc);                   
xlabel('time/day');ylabel('Amplitude');title('remove DC');


subplot(313);
plot(day, kline_NoDcTg_r0);    
xlabel('time/day');ylabel('Amplitude');title('R set to 0');


%imwrite(gcf, 'test.png');   % Use gcf to save the current figure
%print(gcf, '-dpng','d.png');
% saveas(figure_handle,filename,fileformat)
saveas(gcf, '../data/kline/fig1.png',  'png')




%%%%%%%%%%%%%%%%%% 脉冲测试信号 %%%%%%%%%%%%%%%
% Preparing the Desired Impulse Response
%sig = zeros(1,Nsig);        
%sig(1:period:Nsig) = ones(size(1:period:Nsig));


%=================== 输入信号 接入： ====================
Nsig = L_Kline   %1024     

%xm = s(1: Nsig)   % (1:period:Nsig)          % 正弦信号接入
%xm = kline_noDc                              % 只去除DC，不去除TG
%xm = kline_NoDcTg                            % 去除DC，去除TG
xm = kline_NoDcTg_r0                          % 右侧置零


%单窗绘制输入信号
figure(2)
subplot(111);
t_Nsig = linspace(0,  Nsig, Nsig);
plot(t_Nsig, xm );
xlabel('time');ylabel('Amplitude');title('signal before filter');
%%% axis([0, 0.4, -3, 3] );

saveas(gcf, '../data/kline/fig2.png',  'png')

%===================== 输入信号 加窗 =======================
% 输入信号时域加窗
Len_TD_W = Nsig*Len_W_TD_L_Ratio
w_td_full = hamming(Len_TD_W*2);            % flattopwin  kaiser    blackman   hamming  hanning chebwin  bartlett
w_td_half   = w_td_full(1:Len_TD_W)
epsilon = .0001                      % w_td修正，避免 0 / 0
w_td = [w_td_half'  ones(1, Nsig-length(w_td_half))] + epsilon

xm_w = xm .* w_td   % 时域加窗
%xm_w = xm             % 时域不加窗

%---计算修正窗---
w_td_rvs = (1 ./ w_td)             % 加窗后反向修正系数
w_td_rvs_zp = [w_td_rvs   ones(1, Nfft-length(w_td_rvs))]   % 信号点数小于nfft，补零

%----- 绘制 加窗后输入信号 -----
f = linspace(0, fs/2, Nfft_fd/2);  % 一半 linspace(0, pi, 1000):产生0-pi间1000点行矢量

figure(3)
subplot(311);
t_sig = 1:Nsig;          % 定义时间范围和步长
plot(t_sig, xm_w);        % 加窗后、滤波前的信号图像
xlabel('time/sec');ylabel('Amplitude');title('after time-domain windowing');



%--- 加窗后 信号频谱分析 ---
Fs_w  = fft(xm_w, Nfft_fd);
AFs_w = abs(Fs_w);                % 将信号变换到频域及信号频域图的幅值
subplot(312)
plot(f, AFs_w(1 : Nfft_fd/2));   %  滤波前的信号频域图
xlabel('freq/Hz');ylabel('Amplitude');title('Freq-domain before filter');
axis([0, 200, -50, 500000] );

subplot(313)
plot(f, AFs_w(1 : Nfft_fd/2));   %  滤波前的信号频域图
axis([0, 150, -50, 150000] );


saveas(gcf, '../data/kline/fig3.png',  'png')

%---------------- 构建H(k) / 滤波 -------------- 
xm_w_zp = [xm_w zeros(1, Nfft-length(xm_w))];   % zero pad the signal
P = fft(xm_w_zp, Nfft);

%  window it:
%先拷贝P头尾进来到P1(即加矩形窗)
for k = 1 : Nsig
  P1(k) = real( P(k)) + 1i * imag(P(k) );     %error: P(4097): out of bound 4096 (dimensions are 1x4096)
endfor
P1;

% 尾巴高频部分置零
for k = fc : Nfft    % Nsig                        
  P1(k) = 0;
endfor
P1;

%---------------- 频域加窗 & 点乘滤波 -----------------------
Len_W_FD   = fc

%Len_W_FD_R   = fc * Len_W_FD_R_Ratio       
Len_W_FD_R   = ceil( fc * Len_W_FD_R_Ratio      )       
Len_W_FD_BOX = fc*(1- Len_W_FD_R_Ratio) 
w_fd_r_org = hamming(Len_W_FD_R * 2) ; 
w_fd_r     = w_fd_r_org'(Len_W_FD_R+1 :  Len_W_FD_R*2);
w_fd_box   = boxcar(round(Len_W_FD_BOX));       % 四舍五入：round;向上取整：ceil;向下:floor
%w_fd = boxcar(Len_W_FD)     % rectwin  hamming  kaiser %  FIR filter design by window method
w_fd  = [w_fd_box'  w_fd_r] 
w_fd_zp =  [w_fd  zeros(1, Nfft-Len_W_FD)];
%w_fd_zp =  [w_fd  zeros(1, Nfft-Len_W_FD-1)];


%%%%%% 点乘滤波 %%%%%%%

Ym = w_fd_zp .* P1;       % window the fft lowpass data in FD 

%hideal = sin( 2* pi* fc * nfilt/fs) ./ (pi * nfilt)
%hideal = ones( )

%hzp = [h zeros(1, Nfft-L)]   % zero-pad h to FFT size
%H = fft(hzp)               % filter frequency response

%-- 滤波 ---
%y = zeros(1, Nsig + Nfft)   % allocate output+'ringing' vector
%Xm = fft(xmzp)
%Ym = Xm .* H                   % freq domain multiplication

%-------------绘制P频率域波形-------------------
f_fft_full = linspace(0, 2*pi,   Nfft   );
f_fft      = linspace(0, pi,     Nfft/2 );

figure(4)
subplot(411)
plot(f_fft_full, P );  
xlabel('freq/Hz');ylabel('Amplitude'); title('P');
axis([0, 0.5, -200000, 600000] );

%-绘制截取后P频率域波形--
subplot(412)
plot(f_fft_full, P1 );   
xlabel('freq/Hz');ylabel('Amplitude'); title('P1');
axis([0, 0.5, -200000, 600000] );

%---- 窗形状 -----
subplot(413)
t_fc_win = 1:Len_W_FD;
plot(t_fc_win,  w_fd );   
xlabel('freq/Hz');ylabel('Amplitude');title('Ym');
axis([0, 0.5, -20000, 60000] );

%---- P加窗后频率域波形 -----
subplot(414)
plot(f_fft_full, Ym );   
xlabel('freq/Hz');ylabel('Amplitude');title('Ym');
axis([0, 0.5, -20000, 60000] );

saveas(gcf, '../data/kline/fig4.png',  'png')


%------------------ iFFT  -------------------------
ym = real(ifft(Ym));         % inverse transform %频域-> 时域 -> 取real
    %outindex = m*R+1:(m*R+Nfft);    % index prepare for overlap add
    %y(outindex) = y(outindex) + ym; % overlap add
%end

%-------------- 加窗后 幅度修正 -------------
%----时域窗修正-----
%--幅度修正
ym_rvs_amp = C_wc_mul  * ym;    % * C_wc_FD_mul
%--左时域窗波形修正
ym_rvs_TD_L =  ym_rvs_amp .* w_td_rvs_zp;

%----频域窗幅度修正


%---- 最后输出---
k_prc = ym_rvs_amp(1:Nsig)



%绘制滤波前信号
figure(5)
subplot(311);
t_Nsig = linspace(0,  Nsig, Nsig);
plot(t_Nsig, xm );
xlabel('time');ylabel('Amplitude');title('signal before filter');
#axis([0, 0.4, -3, 3] );

#绘制时域加窗后(滤波前)信号
subplot(312);
plot(t_Nsig, xm_w );
xlabel('time');ylabel('Amplitude');title('signal after windowing TD FFT');
#axis([0, 0.4, -3, 3] );

%绘制滤波后信号
subplot(313);
plot(t_Nsig, ym(1:Nsig) );
xlabel('time');ylabel('Amplitude');title('After filter');

saveas(gcf, '../data/kline/fig5.png',  'png')

%-----------单窗绘制最终信号---------------
figure(6)
plot(t_Nsig, ym_rvs_amp(1:Nsig), 'r' );
%hold on;
%plot(t_Nsig, xm, 'b' );
xlabel('time');ylabel('Amplitude');title('After ReviseAmp');

saveas(gcf, '../data/kline/fig6.png',  'png')

%figure(7)
%plot(t_Nsig, ym_rvs_TD_L(1:Nsig), 'r' );
%xlabel('time');ylabel('Amplitude');title('After ReviseTDwin');

endfunction
