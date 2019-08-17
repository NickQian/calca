% 2019.4.2   
% though I use the "Weighter" in the name "CW" but there is no gravity really.
% Only the people's "PB/PE brings profit" thought pull the M back.  
function cwode

tspan = [0, 54];       % 36 times/1. 1year=5.  25=1Period. 250=10Period
%-------------------- Init ------------------------------
%y0 = [pi; 0];                %  init value. [y(t0);  y'=f(t,y) ]   
%y0 = [ 0.123;  -pi];         %  static: [y(t0);  y'(t0);     ]  
y0 = [3.16; 0.24];           %  [ y(t0):3.14+V*t ;    ;   ]  

% [K = 1.115; J = 1.90; V =0.56 ] 
% [K = 1.274; J = 4.012; V =0.0667; y0=[0.123, -pi] ]
M0 = 1.000;
K0 = 1.274000;            % K: W coefficient.  4/*3.14
J0 = 4.012000;             % J: Casine coefficient.
V = 0.2133;        % (PE =15,V=0.0667 ) (PE = 30, V=0.0333)
cnt = 0

%[t,y] = ode15s(@cwe,tspan,y0,options);
%ode = @(t,y) vanderpoldemo(t,y,Mu);
%[t,y] = ode45(ode, tspan, y0);

%----------------- Apply ODE resolver ------------------------
% output: a column vector of t & solution array y(1st column: y1, 2nd: y2
[t,y] = ode45(@cwe, tspan, y0);
%[t,y,te,ye,ie] = ode45(@f,tspan,y0,options);
%fprintf('y is %f \n', y);

% Plot of the solution
plot(t, y(:,1),'-b',  t, y(:,2), '-.r' )
figure;
plot(  y(:,1), y(:,2));
xlabel('t')
ylabel('solution y')
title('CW Equation, \mu = 1')

%axis([tspan(1) tspan(end) -2.5 2.5]); 
%legend('y_1','y_2')


  %--------------------- ODE ---------------------------
  function dydt = cwe_static(t,y)   % no grow
  % Defines the equation for cw.
    dydt = [y(2); J0 * sin(y(1))- K0 * y(1)];      
  end

  function dydt = cwe(t,y)
    coelist = make_coe(M0, K0, J0, t);
    coeJ = coelist(1);
    coeK = coelist(2);
    if isnan(coeJ)     coeJ = J0;  end
    if isnan(coeK)     coeK = K0;  end
    cnt = cnt + 1;
    fprintf('coeJ:%f, coeK: %f, t:%f,V*t:%f, cnt:%f \n', ...
                 coeJ,     coeK,     t,    V*t,   cnt);
    % Defines the equation for cw.
    dydt = [y(2);  coeJ *sin(y(1)-(pi+V*t)) - coeK *(y(1)-(pi+V*t)) ];
    %dydt = [y(2); (1-y(1)^2)*y(2)+ J*sin(y(1))- K*y(1)]; 
    
    fprintf('@2: [%f, %f]  y(1):%f y(2): %f \n',y, y(1),y(2));
  end

  function coes = make_coe(M0, K0, J0, t)
      M = M0 * t;     % ??
      K = K0 * t;
      J = J0 * t;
      Jcoe = J/M;
      Kcoe = K/M;
      coes = [Jcoe, Kcoe];
      %fprintf('t:%f, J:%f, K:%f, M:%f; V*t:%f \n', t, J, K, M,  V*t);
  end

end  % end cwode

