% 加改进 ： 半窗


% 问题1： 直接hg为1是怎么回事?
% 问题2: 无延迟怎么搞定的?

%======================  参数  =========================
pkg load signal

Nsig_min = 1000                  % 至少分析300个交易日
eps = 5.0                    % 容忍误差, 容忍误差内找到点位加半窗。 找不到报错
%Nsig = 1024     % 150  500     % Nfft should > Nsig ? % signal length in samples
fc = 70                         % 50/60/70    X:100   lowpass cutoff frequency in Hz
fs = 4000                       % sampling rate in Hz 

L =  2047 %1023                       %31  511 % FIR filter length in taps
%FFT processing parameters:
M = L                           % nominal window length
Nfft = 2^(ceil(log2(M+L-1)))    % FFT Length. L=31-> Nfft=64

T  = 1;   dt = 1.0/fs;  Nfft_fd = 16384   %8192    % 4096
t = linspace(0, T,    fs);         % 生成2048个点。这些点的间距为 (x2-x1)/(n-1)。
f = linspace(0, fs/2, Nfft_fd/2);  % 一半 linspace(0, pi, 1000):产生0-pi间1000点行矢量
f_fft_full = linspace(0, 2*pi,   Nfft   );
f_fft      = linspace(0, pi,     Nfft/2 );
%f = linspace(0, fs/2, Nfft_fd)  % 一半 linspace(0, pi, 1000):产生0-pi间1000点行矢量

%------加窗 频域幅值修正系数-----
C_wc = 0.5396      % hanning-0.4995  hamming-0.5396   blackman-0.4196
C_wc_mul = 1/C_wc


%========================== 输入信号 ===========================
%%%%%%%%%%%%% 导入数据 %%%%%%%%%%%%%%%%%%%%%%
%[x1, x2, x3, x4] = textread('kline.txt', '%2s  %.3f  %.3f  %.3f')
% fscanf 和 textscan
d = importdata('kline.txt');
kline_all = d';         %d.data  d.textdata  d.rowheaders
k_lastDay = kline_all(length(kline_all))
Nsig_max = length(kline_all)                 


%------倒序找相同点位 ----------

kline_search = []
kline_flp    = []


kline_all_flp = flip(kline_all);        % 先反转
for k =  Nsig_min+1 : Nsig_max   
  head_tail = kline_all_flp(k) -  k_lastDay;
  k
  if abs(head_tail) > eps
       kline_search(k-Nsig_min) = kline_all_flp(k);
  else
       kline_flp = [kline_all_flp(1:Nsig_min) kline_search  kline_all_flp(k)]  %kline_all_flp(k)是最后一个数
       break               % 终止拷贝
  endif
endfor

kline = flip(kline_flp)                % 反转回来

L_Kline = length(kline)

%-------- DC Remove ----------------
dc_sum = 0
for k = 1 : L_Kline
  dc_sum = dc_sum + kline(k)    
endfor
dc = dc_sum/L_Kline

kline_noDc = []
for k = 1:L_Kline
  kline_noDc(k) = kline(k) - dc;
endfor

kline_noDc


%----------------- TG(tree growth) Removal -------------------
% 选定两个bot，找出斜率
bot1 = kline_noDc(1)
bot2 = kline_noDc(L_Kline)
TG_high = (bot2 + bot1)/2


% 画出TG
% tg = (-(TG_high -1)/2 : (TG_high -1)/2)
tg_pre = linspace(bot1,  bot2,   L_Kline   )
tg     = tg_pre - TG_high

% TG remove
%kline_NoDcTg = kline_noDc - tg   % 移除TG
kline_NoDcTg = kline_noDc        % 不移除TG

%--------------------- 右侧移零 ------------------------------
kline_NoDcTg_r0 = kline_NoDcTg - kline_NoDcTg(length(kline_NoDcTg))


%-------------------- 预处理结果绘图 ---------------------------
% 绘图读入数据
figure(1)
day = 1:1:L_Kline

subplot(311);
plot(day, kline);                   
xlabel('time/day');ylabel('Amplitude');title('kline');

subplot(312);
plot(day, kline_noDc);                   
xlabel('time/day');ylabel('Amplitude');title('remove DC');


subplot(313);
plot(day, kline_NoDcTg_r0);    
xlabel('time/day');ylabel('Amplitude');title('R set to 0');

%%%%%%%%%%%%%%%%%% 脉冲测试信号 %%%%%%%%%%%%%%%
%sig = zeros(1,Nsig);
%sig(1:period:Nsig) = ones(size(1:period:Nsig));


%============== 输入信号 接入： ==================
Nsig = L_Kline   %1024     

%xm = s(1: Nsig)   % (1:period:Nsig)          % 正弦信号接入
%xm = kline_noDc                              % 只去除DC，不去除TG
%xm = kline_NoDcTg                            % 去除TG
xm = kline_NoDcTg_r0                          % 右侧置零


%单窗绘制输入信号
figure(2)
subplot(111);
t_Nsig = linspace(0,  Nsig, Nsig)
plot(t_Nsig, xm );
xlabel('time');ylabel('Amplitude');title('signal before filter');
#axis([0, 0.4, -3, 3] );

%================= 输入信号处理 =====================
% 输入信号时域加窗
w_td = hamming(Nsig)            % kaiser    blackman   hamming  hanning chebwin  bartlett
epsilon = .0001                 % avoids 0 / 0
w_td = w_td + epsilon
w_td_rvs = (1 ./ w_td)'         % 加窗后反向修正系数
w_td_rvs_zp = [w_td_rvs   zeros(1, Nfft-length(w_td_rvs))]   % 信号点数小于nfft，补零

%xm_w = xm .* w_td'   % 时域加窗
xm_w = xm             % 时域不加窗


%--- 绘制 加窗后输入信号 ---
figure(3)
subplot(311);
t_sig = 1:Nsig          % 定义时间范围和步长
plot(t_sig, xm_w);        % 加窗后、滤波前的信号图像
xlabel('time/sec');ylabel('Amplitude');title('after time-domain windowing');


%--- 加窗后 信号频谱分析 ---
Fs_w  = fft(xm_w, Nfft_fd);
AFs_w = abs(Fs_w)                % 将信号变换到频域及信号频域图的幅值
subplot(312)
plot(f, AFs_w(1 : Nfft_fd/2));   % @@ 滤波前的信号频域图
xlabel('freq/Hz');ylabel('Amplitude');title('Freq-domain before filter');
axis([0, 200, -50, 500000] );

subplot(313)
plot(f, AFs_w(1 : Nfft_fd/2));   % @@ 滤波前的信号频域图
axis([0, 150, -50, 150000] );

%---------------- 构建H(k) / 滤波 ------------ Preparing the Desired Impulse Response

xm_w_zp = [xm_w zeros(1, Nfft-length(xm_w))]   % zero pad the signal
P = fft(xm_w_zp, Nfft)

%  window it:
%先拷贝P头尾进来到P1(即加矩形窗)
for k = 1 : Nsig
  P1(k) = real( P(k)) + 1i * imag(P(k) );
endfor
P1

% 尾巴高频部分置零
for k = fc : Nfft    % Nsig                        
  P1(k) = 0;
endfor
P1

%hideal = sin( 2* pi* fc * nfilt/fs) ./ (pi * nfilt)
%hideal = ones( )

%---------------- 频域加窗 & 点乘滤波 -----------------------
Len_W_FD = fc
w_fd = boxcar(Len_W_FD)     % rectwin  hamming  kaiser %  FIR filter design by window method
w_fd_zp =  [w_fd'  zeros(1, Nfft-Len_W_FD)]
Ym = w_fd_zp .* P1       % window the fft lowpass data in FD 


%hzp = [h zeros(1, Nfft-L)]   % zero-pad h to FFT size
%H = fft(hzp)               % filter frequency response

%-- 滤波 ---
%y = zeros(1, Nsig + Nfft)   % allocate output+'ringing' vector

%Xm = fft(xmzp)
%Ym = Xm .* H                   % freq domain multiplication


%-------------绘制P频率域波形-------------------
figure(4)
subplot(311)
plot(f_fft_full, P );  
xlabel('freq/Hz');ylabel('Amplitude'); title('P');
axis([0, 0.5, -200000, 600000] );

%-绘制截取后P频率域波形--
subplot(312)
plot(f_fft_full, P1 );   
xlabel('freq/Hz');ylabel('Amplitude'); title('P1');
axis([0, 0.5, -200000, 600000] );

%---- P加窗后频率域波形 -----
subplot(313)

plot(f_fft_full, Ym );   
xlabel('freq/Hz');ylabel('Amplitude');title('Ym');
axis([0, 0.5, -20000, 60000] );


%------------------ iFFT  -------------------------
ym = real(ifft(Ym))         % inverse transform %频域-> 时域 -> 取real
    %outindex = m*R+1:(m*R+Nfft);    % index prepare for overlap add
    %y(outindex) = y(outindex) + ym; % overlap add
%end

%-------------- 加窗后 幅度修正 -------------
ym_amp_rvs = C_wc_mul * ym;

% 正弦校准信号的加窗幅度修正。..利用正弦校准信号频谱的特点
ym_rvs =  ym_amp_rvs .* w_td_rvs_zp



%绘制滤波前信号
figure(5)
subplot(311);
t_Nsig = linspace(0,  Nsig, Nsig)
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

%-----------单窗绘制最终信号---------------
figure(6)
subplot(111);
plot(t_Nsig, ym_amp_rvs(1:Nsig), 'r' );
xlabel('time');ylabel('Amplitude');title('After ReviseAmp_TDwin');
