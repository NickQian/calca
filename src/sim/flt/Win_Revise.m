% window revise - 2021.6.9
M    = 1000
Nsig = M

t=0:1:Nsig-1

w     = hamming(M) 
w_rvs = 1 ./ w

sig = ones(1, Nsig)
sig_w = w' .* sig
sig_flted = 0.47 .* sig_w

sig_rvs = sig_flted .* w_rvs'




figure(1)
subplot(411)
plot(t, sig_w );  

subplot(412)
plot(t, sig_flted );  

subplot(413)
plot(t, w_rvs );  

subplot(414)
plot(t, sig_rvs );  