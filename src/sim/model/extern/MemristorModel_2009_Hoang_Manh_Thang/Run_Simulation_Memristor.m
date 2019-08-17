% This work is based on two papers as follows:
% [1] Dmitri B. Strukov, Gregory S. Snider, Duncan R. Steward and R.
% Stanley Williams, "The missing memristor found", Nature, vol.453,
% pp.80-83, 2008
% [2]Yogesh N Joglekar and Stephen J Wolf, "The elusive memristor: properties of basic electrical circuits", European Journal of Physics, vol. 30, pp. 661-675, 2009. 


% Following is predefined, and you can change them and find the validation
% according to those given in above papers
% To run the simulation, you only need to click on Run after choosing
% values of parameters
% If you have any comment on this, please send directly to Thang Manh Hoang at
% hmthang01@yahoo.com, http://fet.hut.edu.vn/SIP-LAB

% v=v_0 sin(omega_0t)
% D = 10nm = 10e-9 m
% Muy_D = 1e-10 cm2s-1V-1 = 1e-14 m-2s-1V-1
% i_0 = v_0/R_ON=10 mA = 10^-2 A ===> R_ON = 100 ohm
% R_OFF/R_ON = 160 ==> R_OFF = 16e+3 ohm
% w_0/D = 0.5 ==> w_0 = 5nm = 5e-9 m

% The voltage source applied on the Memristor

v_0=1;
bias=0;
omega=1; % frequency (rad/s)
phase=0;
Nuy=-1;
Muy_D=1e-14;
R_OFF=16e+3; % resistance of undoped region
R_ON=100;
D = 10e-9; % Width of Memristor
w_0 = 5e-009; % Width of doped region 
 
f1=figure(1);
hold on;
grid on;

f2=figure(2);
set(0,'CurrentFigure',2)

hold on;
grid on;


for omega=1:5
        comm=sprintf('set_param(''i_v_flux_of_memristor/Sine_wave'',''Amplitude'',''%.30f'')',v_0); % Amplitude
                        eval(comm);
        comm=sprintf('set_param(''i_v_flux_of_memristor/Sine_wave'',''bias'',''%.30f'')',bias); % bias
                        eval(comm);
        comm=sprintf('set_param(''i_v_flux_of_memristor/Sine_wave'',''Frequency'',''%.30f'')',omega); % Frequency
                        eval(comm);
        comm=sprintf('set_param(''i_v_flux_of_memristor/Sine_wave'',''phase'',''%.30f'')',phase); % Phase
                        eval(comm);

       

             

         % Other parameters of equation (7) in the page 665 of [2]


         comm=sprintf('set_param(''i_v_flux_of_memristor/i_v_flux/Nuy'',''value'',''%.30f'')',Nuy); % Nuy
                        eval(comm);     
         comm=sprintf('set_param(''i_v_flux_of_memristor/i_v_flux/Muy_D'',''value'',''%.30f'')',Muy_D); % Muy_D
                        eval(comm);  
         comm=sprintf('set_param(''i_v_flux_of_memristor/i_v_flux/R_OFF'',''value'',''%.30f'')',R_OFF); % Resistance of undoped
                        eval(comm);                
         comm=sprintf('set_param(''i_v_flux_of_memristor/i_v_flux/R_ON'',''value'',''%.30f'')',R_ON); % Resistance of doped
                        eval(comm);
         comm=sprintf('set_param(''i_v_flux_of_memristor/i_v_flux/D'',''value'',''%.30f'')',D); % Width of Memristor
                        eval(comm);   
         comm=sprintf('set_param(''i_v_flux_of_memristor/i_v_flux/w_0'',''value'',''%.30f'')',w_0); % Width of doped region
                        eval(comm);   
         sim('i_v_flux_of_memristor');        


         disp('-----------------------------------------Results-----------------------------------------');
        t_0=D*D/Muy_D/v_0;
         comm=sprintf('disp(''t_0=D^2/Muy_D/v_0 = %e'')',t_0); % Time that the dopants need to travel distance D under constant voltage v_0, chosen app. 1ms
                        eval(comm);   
        comm=sprintf('disp(''Omega_0=2*pi/t_0 = %e'')',2*pi/t_0); % Frequency scales for memristor circuit, chosen typically as given in [1]
                        eval(comm);  
        comm=sprintf('disp(''Q_0 = %e'')',Q_0); % 
                        eval(comm);


        qmax = Q_0*(1-w_0/D);

        comm=sprintf('disp(''q_max(t) = Q_0(1-w_0/D) = %e'')',qmax);
                        eval(comm);  

% plot of v-i 
        set(0,'CurrentFigure',1)
        f1=plot(v,i);
        set(0,'CurrentFigure',2)
        f2=plot(flux,q);
end 

 
                