#! /bin/octave  -qf

printf ("Hello, octave!\n"); 

%======================   ===================================
pkg load signal

LEN_FIX = 1150;   % sh_gem: 1500;   sug: 850


% FC  sh & sz: 70(for buy)-81(for sale)      gem: 55  star: ?
%      sug: 90
[k_org, k_prc] = fsf_zph_ltwin('../data/kline/sh_K.txt', fc_in=70, LEN_FIX)   %1450  %  fn, fc_in, le$



%-------- 
%[tp, pt] = trav(k_org, k_prc);


%--------------- plot -----------------

