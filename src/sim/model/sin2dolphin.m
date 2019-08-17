function sin2dolphin

coeJ    = 2.77102;
coeJ_F  = 0.8;
coeK    = 0.9805170;
coeK_F  = 0.7;
coeG    = 5.10;
G_F     = 0.3;               % horizontal flexibility of G

%------- dolphin shape -----
B   = 0.4235;   % pi - asin(A); % 0.5235;
A   = -sin(-pi + B);   %0.2

fprintf('B:%f, A: %f  \n',     B, A);

%-----define plot area-----
x = -1.3*pi: 0.05 : 1.3*pi;

%---------------- curve --------------------
y = coeJ*sin(x);

y_up = coeJ*(sin(coeJ_F*x) + A);
y_upleft = coeJ*(sin(coeJ_F*x+B) + A);

y_down = coeJ*(sin(coeJ_F*x) - A);
y_downright = coeJ*(sin(coeJ_F*x-B) - A);

%y_val = -coeK * x;
y_val = -coeK *(sin(coeK_F*x));

y_grow = coeG*(-G_F*x).^5;

%--------------------------------
plot(x,y, 'y',     x,[pi], x,[-pi], x, 0,   ...
     [-pi -pi],[-6,6], ':g',   x, y_upleft, 'r', [pi  pi],[-6,6], ':c', x, y_downright, 'm',  ...
     x, y_val, 'b', x, y_grow, 'k');

xlabel('y');
ylabel('money');
title('dolphin shape, \mu = 1');

end 