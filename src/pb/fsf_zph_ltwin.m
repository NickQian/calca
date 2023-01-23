% ʱ����߼Ӵ����̶���������750�㡣
% 1) ����Ӧ�����й�Ʊ��eg. ���״���ֻ��800��������
% 2) �����ڲ�ָͬ����ʱ��Ա�(ʱ������ͬ��)

%������ֻ���� 1)

%pkg load signal

%======================  ����  ===================================
function [k_org, k_prc] = fsf_zph_ltwin(fn, fc_in, len_in)   %'kline.txt'


Nsig_FixLen = len_in  %len_in                      % ����3��       
fc = fc_in %fc_in                         % 50/60/70    X:100   lowpass cutoff frequency in Hz
Len_W_FD_R_Ratio = 1/7             % Ƶ�������Ҳ�Ӵ�����,Ϊ����box��1/Len_W_FD_R_DIV
Len_W_TD_L_Ratio = 1/10            % �೤���������ݼӴ�


%FFT processing parameters:
L =  2047;  M=L         % FIR filter length in taps  % nominal window length 
Nfft = 2^(ceil(log2(M+L-1)))    % FFT Length. L=31-> Nfft=64

fs = 4000                       % sampling rate in Hz 
T  = 1;   dt = 1.0/fs;  
Nfft_fd = 16384   %8192    % 4096

%t = linspace(0, T,    fs);         % ����2048���㡣��Щ��ļ��Ϊ (x2-x1)/(n-1)��


%---- ʱ�� ��ֵ����ϵ��-----
C_wc = 0.5396      % hanning-0.4995  hamming-0.5396   blackman-0.4196
C_wc_mul = 1/C_wc

%---- Ƶ�� ��ֵ����ϵ��-----
C_wc_FD = 1.313
C_wc_FD_mul = 1/C_wc_FD;

%========================== �����ź� ==============================
%----- �������� ------
d  = importdata(fn);    %(fn)   ("kline_sh.txt") % fscanf �� textscan       %[x1, x2, x3, x4] = textread('kline.txt', '%2s  %.3f  %.3f  %.3f')
kline_all  = d';         %d.data  d.textdata  d.rowheaders
k1_lastDay = kline_all(length(kline_all))   
Nsig_max   = length(kline_all)                


%------------- ��ȡ�Ҳ�Nsig_FixLen�����ź� -----------
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
% ѡ������bot���ҳ�б��
bot1 = kline_noDc(1)
bot2 = kline_noDc(L_Kline)
TG_high = (bot2 + bot1)/2;

% ����TG
% tg = (-(TG_high -1)/2 : (TG_high -1)/2)
tg_pre = linspace(bot1,  bot2,   L_Kline   );
tg     = tg_pre - TG_high;

% TG remove
%kline_NoDcTg = kline_noDc - tg   % �Ƴ�TG
kline_NoDcTg = kline_noDc;        % ���Ƴ�TG

%--------------------- �Ҳ����� ------------------------------
kline_NoDcTg_r0 = kline_NoDcTg - kline_NoDcTg(length(kline_NoDcTg))


%-------------------- Ԥ��������ͼ ---------------------------
% ��ͼ��������
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




%%%%%%%%%%%%%%%%%% ��������ź� %%%%%%%%%%%%%%%
% Preparing the Desired Impulse Response
%sig = zeros(1,Nsig);        
%sig(1:period:Nsig) = ones(size(1:period:Nsig));


%=================== �����ź� ���룺 ====================
Nsig = L_Kline   %1024     

%xm = s(1: Nsig)   % (1:period:Nsig)          % �����źŽ���
%xm = kline_noDc                              % ֻȥ��DC����ȥ��TG
%xm = kline_NoDcTg                            % ȥ��DC��ȥ��TG
xm = kline_NoDcTg_r0                          % �Ҳ�����


%�������������ź�
figure(2)
subplot(111);
t_Nsig = linspace(0,  Nsig, Nsig);
plot(t_Nsig, xm );
xlabel('time');ylabel('Amplitude');title('signal before filter');
%%% axis([0, 0.4, -3, 3] );

saveas(gcf, '../data/kline/fig2.png',  'png')

%===================== �����ź� �Ӵ� =======================
% �����ź�ʱ��Ӵ�
Len_TD_W = Nsig*Len_W_TD_L_Ratio
w_td_full = hamming(Len_TD_W*2);            % flattopwin  kaiser    blackman   hamming  hanning chebwin  bartlett
w_td_half   = w_td_full(1:Len_TD_W)
epsilon = .0001                      % w_td���������� 0 / 0
w_td = [w_td_half'  ones(1, Nsig-length(w_td_half))] + epsilon

xm_w = xm .* w_td   % ʱ��Ӵ�
%xm_w = xm             % ʱ�򲻼Ӵ�

%---����������---
w_td_rvs = (1 ./ w_td)             % �Ӵ���������ϵ��
w_td_rvs_zp = [w_td_rvs   ones(1, Nfft-length(w_td_rvs))]   % �źŵ���С��nfft������

%----- ���� �Ӵ��������ź� -----
f = linspace(0, fs/2, Nfft_fd/2);  % һ�� linspace(0, pi, 1000):����0-pi��1000����ʸ��

figure(3)
subplot(311);
t_sig = 1:Nsig;          % ����ʱ�䷶Χ�Ͳ���
plot(t_sig, xm_w);        % �Ӵ����˲�ǰ���ź�ͼ��
xlabel('time/sec');ylabel('Amplitude');title('after time-domain windowing');



%--- �Ӵ��� �ź�Ƶ�׷��� ---
Fs_w  = fft(xm_w, Nfft_fd);
AFs_w = abs(Fs_w);                % ���źű任��Ƶ���ź�Ƶ��ͼ�ķ�ֵ
subplot(312)
plot(f, AFs_w(1 : Nfft_fd/2));   %  �˲�ǰ���ź�Ƶ��ͼ
xlabel('freq/Hz');ylabel('Amplitude');title('Freq-domain before filter');
axis([0, 200, -50, 500000] );

subplot(313)
plot(f, AFs_w(1 : Nfft_fd/2));   %  �˲�ǰ���ź�Ƶ��ͼ
axis([0, 150, -50, 150000] );


saveas(gcf, '../data/kline/fig3.png',  'png')

%---------------- ����H(k) / �˲� -------------- 
xm_w_zp = [xm_w zeros(1, Nfft-length(xm_w))];   % zero pad the signal
P = fft(xm_w_zp, Nfft);

%  window it:
%�ȿ���Pͷβ������P1(���Ӿ��δ�)
for k = 1 : Nsig
  P1(k) = real( P(k)) + 1i * imag(P(k) );     %error: P(4097): out of bound 4096 (dimensions are 1x4096)
endfor
P1;

% β�͸�Ƶ��������
for k = fc : Nfft    % Nsig                        
  P1(k) = 0;
endfor
P1;

%---------------- Ƶ��Ӵ� & ����˲� -----------------------
Len_W_FD   = fc

%Len_W_FD_R   = fc * Len_W_FD_R_Ratio       
Len_W_FD_R   = ceil( fc * Len_W_FD_R_Ratio      )       
Len_W_FD_BOX = fc*(1- Len_W_FD_R_Ratio) 
w_fd_r_org = hamming(Len_W_FD_R * 2) ; 
w_fd_r     = w_fd_r_org'(Len_W_FD_R+1 :  Len_W_FD_R*2);
w_fd_box   = boxcar(round(Len_W_FD_BOX));       % �������룺round;����ȡ����ceil;����:floor
%w_fd = boxcar(Len_W_FD)     % rectwin  hamming  kaiser %  FIR filter design by window method
w_fd  = [w_fd_box'  w_fd_r] 
w_fd_zp =  [w_fd  zeros(1, Nfft-Len_W_FD)];
%w_fd_zp =  [w_fd  zeros(1, Nfft-Len_W_FD-1)];


%%%%%% ����˲� %%%%%%%

Ym = w_fd_zp .* P1;       % window the fft lowpass data in FD 

%hideal = sin( 2* pi* fc * nfilt/fs) ./ (pi * nfilt)
%hideal = ones( )

%hzp = [h zeros(1, Nfft-L)]   % zero-pad h to FFT size
%H = fft(hzp)               % filter frequency response

%-- �˲� ---
%y = zeros(1, Nsig + Nfft)   % allocate output+'ringing' vector
%Xm = fft(xmzp)
%Ym = Xm .* H                   % freq domain multiplication

%-------------����PƵ������-------------------
f_fft_full = linspace(0, 2*pi,   Nfft   );
f_fft      = linspace(0, pi,     Nfft/2 );

figure(4)
subplot(411)
plot(f_fft_full, P );  
xlabel('freq/Hz');ylabel('Amplitude'); title('P');
axis([0, 0.5, -200000, 600000] );

%-���ƽ�ȡ��PƵ������--
subplot(412)
plot(f_fft_full, P1 );   
xlabel('freq/Hz');ylabel('Amplitude'); title('P1');
axis([0, 0.5, -200000, 600000] );

%---- ����״ -----
subplot(413)
t_fc_win = 1:Len_W_FD;
plot(t_fc_win,  w_fd );   
xlabel('freq/Hz');ylabel('Amplitude');title('Ym');
axis([0, 0.5, -20000, 60000] );

%---- P�Ӵ���Ƶ������ -----
subplot(414)
plot(f_fft_full, Ym );   
xlabel('freq/Hz');ylabel('Amplitude');title('Ym');
axis([0, 0.5, -20000, 60000] );

saveas(gcf, '../data/kline/fig4.png',  'png')


%------------------ iFFT  -------------------------
ym = real(ifft(Ym));         % inverse transform %Ƶ��-> ʱ�� -> ȡreal
    %outindex = m*R+1:(m*R+Nfft);    % index prepare for overlap add
    %y(outindex) = y(outindex) + ym; % overlap add
%end

%-------------- �Ӵ��� �������� -------------
%----ʱ������-----
%--��������
ym_rvs_amp = C_wc_mul  * ym;    % * C_wc_FD_mul
%--��ʱ�򴰲�������
ym_rvs_TD_L =  ym_rvs_amp .* w_td_rvs_zp;

%----Ƶ�򴰷�������


%---- ������---
k_prc = ym_rvs_amp(1:Nsig)



%�����˲�ǰ�ź�
figure(5)
subplot(311);
t_Nsig = linspace(0,  Nsig, Nsig);
plot(t_Nsig, xm );
xlabel('time');ylabel('Amplitude');title('signal before filter');
#axis([0, 0.4, -3, 3] );

#����ʱ��Ӵ���(�˲�ǰ)�ź�
subplot(312);
plot(t_Nsig, xm_w );
xlabel('time');ylabel('Amplitude');title('signal after windowing TD FFT');
#axis([0, 0.4, -3, 3] );

%�����˲����ź�
subplot(313);
plot(t_Nsig, ym(1:Nsig) );
xlabel('time');ylabel('Amplitude');title('After filter');

saveas(gcf, '../data/kline/fig5.png',  'png')

%-----------�������������ź�---------------
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
