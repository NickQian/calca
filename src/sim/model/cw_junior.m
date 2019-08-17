% simulate the forces on A
% v.01 - cw_junior don't have predictive function 2019.4.2   
%   
function cw_junior

tspan = [0, 15];       % [0, 54]
%-------------------- Init ------------------------------  
%y0 = [0.112; 2.174];            %  [y(t0);  y'(t0)]; 2.13-> 2.174 -> 2.18. 
y0 = [6.166; -0.4820];           % -=>;  =->

%-------------------- Coefficients ------------------------------
% [K = 1.115; J = 1.90; V =0.56 ] 
% [K = 1.274; J = 4.012; V =0.0667; y0=[0.123, -pi] ]
M0     = 1.000;
%J0     = 4.0110200;             % J: Casine coefficient. means power of the casine. eg. 4.012
%K0     = 1.2738854;             % K: W coefficient.  1.2738854=4/*3.14.  4 means pb_top=4x1
J0     = 2.871030;
coeJ_F = 0.8;                    % make the position in [-pi, pi]
K0     = 0.950400;               % 0.980500 -> 0.980600
coeK_F = 0.8;
G0     = 5.10;
coeG_F = 0.30;
R0     = 0.1;                    % 0.1 - 0.3

%--- dolphin parameters ---
% dolphin amplifies the amplitude
%Ja_abs = 0.4;                   % max:1=sin(90`). up/down amplify shift absolute value
Jb_abs = 0.315;  
Ja_abs = -sin(-pi + Jb_abs);     % left/right phase shift absolute value

%--- gdp parameter ---
Vgdp   = 0.1355;                 % (PE =15,V=0.0667 ) (PE = 30, V=0.0333)


%%%--- for debug & plot ---
cnt      = 0;
Cmoney   = [];
Kmoney   = [];
Gmoney   = [];
Almoney  = [];
AllForce = [];
HH =       [];
LL =       [];

T_intl   = [];

%----------------- Apply ODE resolver ------------------------
% output: a column vector of t & solution array y(1st column: y1, 2nd: y2
% ode45 options:(RelTol, AbsTol, NormControl),
%                (OutputFcn, OutputSel, Refine, Stats), NonNegative, Events,
%                (MaxStep, InitialStep), (Mass , MStateDependence )
RAND_LIST = make_rand();                         %fprintf('$:%f \n', RAND_LIST);
opts = odeset('RelTol',1e-6,'AbsTol',[1e-6]);    % default: 1e-3
[t,y] = ode45(@cwe_randist, tspan, y0, opts);


%[t,y,te,ye,ie] = ode45(@f,tspan,y0,options);  
% te: time of event, ye: y at event, ie: index of index
%ode = @(t,y) vanderpoldemo(t,y,Mu);

%disp('@@t:') , disp(t);

%-------------- Plot --------------
plot(t, y(:,1),'-k',  t, y(:,2), '-.y',  t, [0], ':b',  t, [pi], '-.y',   ...    % t, [0], ':b',  t, [2*pi], '-.b', 
     T_intl, Cmoney, '-g',    T_intl, Kmoney, '-.m',   ...   %  T_intl, Gmoney/20, '-c',    T_intl, Almoney/20, '-r',
     T_intl, HH, 'b',   T_intl, LL, 'b');
    
figure;
plot(  y(:,1), y(:,2));              %  y vs. y'
%plot(  y(:,1), CMoney);
xlabel('t')
ylabel('solution y')
title('CW Equation, \mu = 1')

%axis([tspan(1) tspan(end) -2.5 2.5]); 
%legend('y_1','y_2')


  %-------------------------- ODE -------------------------------------------
  %----  sin shape(money in-out) raw vibration ----
  % 
  function dydt = cwe_raw(t,y)        % no gdp grow. with J & K
  % Defines the equation for cw.
    dydt = [y(2); J0 * sin(y(1))- K0 * y(1)];      
  end


  %---- raw with pi shift----
  function dydt = cwe_pishft(t,y)     % no policy Resist
    cnt = cnt + 1;    
    fprintf('J:%f, K: %f, t:%f, [%f, %f],  cnt:%f  \n', ...
             J0,    K0,   t,    y(1),y(2), cnt);
    Cmoney(cnt)  = J0 *sin(y(1)-pi);
    Kmoney(cnt)  = -K0 *(y(1)-pi);
    Almoney(cnt) = Cmoney(cnt) + Kmoney(cnt); %J0 * sin(y(1)-pi ) - K0 *(y(1)-pi);
    T_intl(cnt)  = t;
    % Defines the equation for cw.
    dydt = [y(2);  J0 * sin(y(1)-pi ) - K0 *(y(1)-pi)];
  end


  %---- with "dolphin shape" ----
  % also with "dolphin shape" money in-out, no gdp grow
  function dydt = cwe_dolphin(t,y)   % no policy Resist
    coelist = make_coe(M0, K0, J0, t, y);
    coeJ = coelist(1);
    coeK = coelist(2);
    Ja = coelist(3);
    Jb = coelist(4);
    if isnan(coeJ)     coeJ = J0;  end       % process NaN
    if isnan(coeK)     coeK = K0;  end
    cnt = cnt + 1;    
    fprintf('coeJ:%f, coeK: %f, Ja:%f, Jb:%f, t:%f, [%f, %f], cnt:%f  \n', ...
             coeJ,    coeK,     Ja,    Jb,    t,    y(1),y(2), cnt);
    Cmoney(cnt) = coeJ*(sin(y(1)-pi +Jb) + Ja);
    Kmoney(cnt)  = -coeK *(y(1)-pi);
    Almoney(cnt) =  Cmoney(cnt) + Kmoney(cnt);
    T_intl(cnt) = t;
    % Defines the equation for cw.
    dydt = [y(2);  coeJ *( sin(y(1)-pi +Jb) + Ja) - coeK *(y(1)-pi) ];
  end


  %---- with gdp grow ----
  % also with "dolphin shape" money in-out and gdp grow, but no resist
  function dydt = cwe_grow(t,y)   % no policy Resist
    coelist = make_coe(M0, K0, G0, J0, t, y);
    coeJ = coelist(1);
    coeK = coelist(2);
    coeG = coelist(3);
    Ja = coelist(4);
    Jb = coelist(5);
    if isnan(coeJ)     coeJ = J0;  end       % process NaN
    if isnan(coeK)     coeK = K0;  end
    if isnan(coeG)     coeG = G0;  end
    cnt = cnt + 1;    
    fprintf('coeJ:%f, coeK: %f, Ja:%f, Jb:%f, t:%f, [%f, %f], cnt:%f  \n', ...
             coeJ,    coeK,     Ja,    Jb,    t,    y(1),y(2), cnt);
    Cmoney(cnt)  = coeJ *( sin(coeJ_F*(y(1)-(pi + Vgdp*t) +Jb)) + Ja);
    Kmoney(cnt)  = -coeK *(sin(coeK_F*(y(1)-(pi + Vgdp*t))) );     %-coeK *(y(1)-(pi + Vgdp*t));
    Gmoney(cnt)  = -coeG*(coeG_F*( y(1)-(pi + Vgdp*t)) ).^5;        
    Almoney(cnt) =  Cmoney(cnt) + Kmoney(cnt) + Gmoney(cnt);
    LL(cnt) = Vgdp*t;
    HH(cnt) = LL(cnt) + 2*pi;
    T_intl(cnt)  = t;
    % Defines the equation for cw.
    dydt = [y(2);  coeJ *( sin(coeJ_F*(y(1)-(pi + Vgdp*t) +Jb)) + Ja) - coeK *(sin(coeK_F*(y(1)-(pi + Vgdp*t))) ) - coeG*(coeG_F*( y(1)-(pi + Vgdp*t)) ).^5 ];
  end


  %---- with "random disturbance" ---- 
  function dydt = cwe_randist(t,y)   % no policy Resist
    coelist = make_coe(M0, K0, G0, J0, t, y);
    coeJ = coelist(1);
    coeK = coelist(2);
    coeG = coelist(3);
    Ja = coelist(4);
    Jb = coelist(5);
    if isnan(coeJ)     coeJ = J0;  end       % process NaN
    if isnan(coeK)     coeK = K0;  end
    if isnan(coeG)     coeG = G0;  end
    cnt = cnt + 1;    
    LL(cnt) = Vgdp*t;
    HH(cnt) = LL(cnt) + 2*pi;
    T_intl(cnt)  = t;
    rand = RAND_LIST(cnt);
    Cmoney(cnt)  = coeJ *( sin(coeJ_F*(y(1)-(pi + Vgdp*t) +Jb)) + Ja);
    Kmoney(cnt)  = -coeK *(sin(coeK_F*(y(1)-(pi + Vgdp*t))) );     %-coeK *(y(1)-(pi + Vgdp*t));
    Gmoney(cnt)  = -coeG*(coeG_F*( y(1)-(pi + Vgdp*t)) ).^5;        
    Almoney(cnt) =  Cmoney(cnt) + Kmoney(cnt) + Gmoney(cnt);
    AllForce(cnt) = Almoney(cnt)+ R0*rand;
    
    fprintf('Cny:%f,Kny:%f,Alny:%f,AlnyR:%f, t:%f, [%f, %f], cnt:%f, rand: %f  \n', ...
             Cmoney(cnt), Kmoney(cnt), Almoney(cnt),AllForce(cnt),t,    y(1),y(2), cnt, rand);
 
    % Defines the equation for cw.
    dydt = [y(2);  coeJ *( sin(coeJ_F*(y(1)-(pi + Vgdp*t) +Jb)) + Ja) - coeK *(sin(coeK_F*(y(1)-(pi + Vgdp*t))) ) - coeG*(coeG_F*( y(1)-(pi + Vgdp*t)) ).^5 + R0*rand   ];  %+ R0*rand    
   end


  %---- with policy Resist ----
  % with "dolphin shape" & gdp grow & resist
  function dydt = cwe_junior(t,y)

    % Defines the equation for cw.
    %dydt = [y(2); (1-y(1)^2)*y(2)+ J*sin(y(1))- K*y(1)];     
    
  end



  %-------------- change J & K during run ----------------------
  function coes = make_coe(M0, K0, G0, J0, t, y)
      M = M0 * t;     % ??
      K = K0 * t;
      G = G0 * t;
      if y(2)>0
          Ja = Ja_abs;                        % a is amplify shift
          Jb = Jb_abs ;                      % b is phase shift
          Gcoe = 1;
      else
          Ja = -Ja_abs;
          Jb = -Jb_abs;          
          Gcoe = G/M;
      end
      J = J0 *  t;
      Jcoe = J/M;
      Kcoe = K/M;
      %Gcoe = G/M;
      coes = [Jcoe, Kcoe, Gcoe, Ja, Jb];
      %fprintf('t:%f, J:%f, K:%f, M:%f; V*t:%f \n', t, J, K, M,  V*t);
  end

  function ranlist = make_rand()
      ranlist = rand(1, 9000000,'double') -0.5 ;
  end

end  % end cwode


