% �ӸĽ� �� �봰


% ����1�� ֱ��hgΪ1����ô����?
% ����2: ���ӳ���ô�㶨��?

%======================  ����  =========================
pkg load signal

Nsig_min = 1000                  % ���ٷ���300��������
eps = 5.0                    % �������, ����������ҵ���λ�Ӱ봰�� �Ҳ�������
%Nsig = 1024     % 150  500     % Nfft should > Nsig ? % signal length in samples
fc = 70                         % 50/60/70    X:100   lowpass cutoff frequency in Hz
fs = 4000                       % sampling rate in Hz 

L =  2047 %1023                       %31  511 % FIR filter length in taps
%FFT processing parameters:
M = L                           % nominal window length
Nfft = 2^(ceil(log2(M+L-1)))    % FFT Length. L=31-> Nfft=64

T  = 1;   dt = 1.0/fs;  Nfft_fd = 16384   %8192    % 4096
t = linspace(0, T,    fs);         % ����2048���㡣��Щ��ļ��Ϊ (x2-x1)/(n-1)��
f = linspace(0, fs/2, Nfft_fd/2);  % һ�� linspace(0, pi, 1000):����0-pi��1000����ʸ��
f_fft_full = linspace(0, 2*pi,   Nfft   );
f_fft      = linspace(0, pi,     Nfft/2 );
%f = linspace(0, fs/2, Nfft_fd)  % һ�� linspace(0, pi, 1000):����0-pi��1000����ʸ��

%------�Ӵ� Ƶ���ֵ����ϵ��-----
C_wc = 0.5396      % hanning-0.4995  hamming-0.5396   blackman-0.4196
C_wc_mul = 1/C_wc


%========================== �����ź� ===========================
%%%%%%%%%%%%% �������� %%%%%%%%%%%%%%%%%%%%%%
%[x1, x2, x3, x4] = textread('kline.txt', '%2s  %.3f  %.3f  %.3f')
% fscanf �� textscan
d = importdata('kline.txt');
kline_all = d';         %d.data  d.textdata  d.rowheaders
k_lastDay = kline_all(length(kline_all))
Nsig_max = length(kline_all)                 


%------��������ͬ��λ ----------

kline_search = []
kline_flp    = []


kline_all_flp = flip(kline_all);        % �ȷ�ת
for k =  Nsig_min+1 : Nsig_max   
  head_tail = kline_all_flp(k) -  k_lastDay;
  k
  if abs(head_tail) > eps
       kline_search(k-Nsig_min) = kline_all_flp(k);
  else
       kline_flp = [kline_all_flp(1:Nsig_min) kline_search  kline_all_flp(k)]  %kline_all_flp(k)�����һ����
       break               % ��ֹ����
  endif
endfor

kline = flip(kline_flp)                % ��ת����

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
% ѡ������bot���ҳ�б��
bot1 = kline_noDc(1)
bot2 = kline_noDc(L_Kline)
TG_high = (bot2 + bot1)/2


% ����TG
% tg = (-(TG_high -1)/2 : (TG_high -1)/2)
tg_pre = linspace(bot1,  bot2,   L_Kline   )
tg     = tg_pre - TG_high

% TG remove
%kline_NoDcTg = kline_noDc - tg   % �Ƴ�TG
kline_NoDcTg = kline_noDc        % ���Ƴ�TG

%--------------------- �Ҳ����� ------------------------------
kline_NoDcTg_r0 = kline_NoDcTg - kline_NoDcTg(length(kline_NoDcTg))


%-------------------- Ԥ��������ͼ ---------------------------
% ��ͼ��������
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

%%%%%%%%%%%%%%%%%% ��������ź� %%%%%%%%%%%%%%%
%sig = zeros(1,Nsig);
%sig(1:period:Nsig) = ones(size(1:period:Nsig));


%============== �����ź� ���룺 ==================
Nsig = L_Kline   %1024     

%xm = s(1: Nsig)   % (1:period:Nsig)          % �����źŽ���
%xm = kline_noDc                              % ֻȥ��DC����ȥ��TG
%xm = kline_NoDcTg                            % ȥ��TG
xm = kline_NoDcTg_r0                          % �Ҳ�����


%�������������ź�
figure(2)
subplot(111);
t_Nsig = linspace(0,  Nsig, Nsig)
plot(t_Nsig, xm );
xlabel('time');ylabel('Amplitude');title('signal before filter');
#axis([0, 0.4, -3, 3] );

%================= �����źŴ��� =====================
% �����ź�ʱ��Ӵ�
w_td = hamming(Nsig)            % kaiser    blackman   hamming  hanning chebwin  bartlett
epsilon = .0001                 % avoids 0 / 0
w_td = w_td + epsilon
w_td_rvs = (1 ./ w_td)'         % �Ӵ���������ϵ��
w_td_rvs_zp = [w_td_rvs   zeros(1, Nfft-length(w_td_rvs))]   % �źŵ���С��nfft������

%xm_w = xm .* w_td'   % ʱ��Ӵ�
xm_w = xm             % ʱ�򲻼Ӵ�


%--- ���� �Ӵ��������ź� ---
figure(3)
subplot(311);
t_sig = 1:Nsig          % ����ʱ�䷶Χ�Ͳ���
plot(t_sig, xm_w);        % �Ӵ����˲�ǰ���ź�ͼ��
xlabel('time/sec');ylabel('Amplitude');title('after time-domain windowing');


%--- �Ӵ��� �ź�Ƶ�׷��� ---
Fs_w  = fft(xm_w, Nfft_fd);
AFs_w = abs(Fs_w)                % ���źű任��Ƶ���ź�Ƶ��ͼ�ķ�ֵ
subplot(312)
plot(f, AFs_w(1 : Nfft_fd/2));   % @@ �˲�ǰ���ź�Ƶ��ͼ
xlabel('freq/Hz');ylabel('Amplitude');title('Freq-domain before filter');
axis([0, 200, -50, 500000] );

subplot(313)
plot(f, AFs_w(1 : Nfft_fd/2));   % @@ �˲�ǰ���ź�Ƶ��ͼ
axis([0, 150, -50, 150000] );

%---------------- ����H(k) / �˲� ------------ Preparing the Desired Impulse Response

xm_w_zp = [xm_w zeros(1, Nfft-length(xm_w))]   % zero pad the signal
P = fft(xm_w_zp, Nfft)

%  window it:
%�ȿ���Pͷβ������P1(���Ӿ��δ�)
for k = 1 : Nsig
  P1(k) = real( P(k)) + 1i * imag(P(k) );
endfor
P1

% β�͸�Ƶ��������
for k = fc : Nfft    % Nsig                        
  P1(k) = 0;
endfor
P1

%hideal = sin( 2* pi* fc * nfilt/fs) ./ (pi * nfilt)
%hideal = ones( )

%---------------- Ƶ��Ӵ� & ����˲� -----------------------
Len_W_FD = fc
w_fd = boxcar(Len_W_FD)     % rectwin  hamming  kaiser %  FIR filter design by window method
w_fd_zp =  [w_fd'  zeros(1, Nfft-Len_W_FD)]
Ym = w_fd_zp .* P1       % window the fft lowpass data in FD 


%hzp = [h zeros(1, Nfft-L)]   % zero-pad h to FFT size
%H = fft(hzp)               % filter frequency response

%-- �˲� ---
%y = zeros(1, Nsig + Nfft)   % allocate output+'ringing' vector

%Xm = fft(xmzp)
%Ym = Xm .* H                   % freq domain multiplication


%-------------����PƵ������-------------------
figure(4)
subplot(311)
plot(f_fft_full, P );  
xlabel('freq/Hz');ylabel('Amplitude'); title('P');
axis([0, 0.5, -200000, 600000] );

%-���ƽ�ȡ��PƵ������--
subplot(312)
plot(f_fft_full, P1 );   
xlabel('freq/Hz');ylabel('Amplitude'); title('P1');
axis([0, 0.5, -200000, 600000] );

%---- P�Ӵ���Ƶ������ -----
subplot(313)

plot(f_fft_full, Ym );   
xlabel('freq/Hz');ylabel('Amplitude');title('Ym');
axis([0, 0.5, -20000, 60000] );


%------------------ iFFT  -------------------------
ym = real(ifft(Ym))         % inverse transform %Ƶ��-> ʱ�� -> ȡreal
    %outindex = m*R+1:(m*R+Nfft);    % index prepare for overlap add
    %y(outindex) = y(outindex) + ym; % overlap add
%end

%-------------- �Ӵ��� �������� -------------
ym_amp_rvs = C_wc_mul * ym;

% ����У׼�źŵļӴ�����������..��������У׼�ź�Ƶ�׵��ص�
ym_rvs =  ym_amp_rvs .* w_td_rvs_zp



%�����˲�ǰ�ź�
figure(5)
subplot(311);
t_Nsig = linspace(0,  Nsig, Nsig)
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

%-----------�������������ź�---------------
figure(6)
subplot(111);
plot(t_Nsig, ym_amp_rvs(1:Nsig), 'r' );
xlabel('time');ylabel('Amplitude');title('After ReviseAmp_TDwin');
