function [s,w]= EntropyWeight(x,ind)
    %??????????(?????????????
    %x???????, ????????, ????????
    %ind????????????????????1???????2??????
    %s???????????w??????
    [n,m]=size(x); % n indicator, m samples
    fprintf('n=%d samples, m = %d indicators \n', n, m )
    %%????????
    for i=1:m        
        if ind(i)==1 %???????

            X(:,i)=guiyi(x(:,i),1, 0.0001, 0.9999);    %?????[0,1], 0????              
        else %???????
            X(:,i)=guiyi(x(:,i),2, 0.0001, 0.9999);
        end        
    end
    fprintf('normalized X is:\n' )
    disp(X)
    
    %%???j??????i??????????p(i,j)
    for i=1:n
        for j=1:m
            p(i,j)=X(i,j)/sum(X(:,j));
            %fprintf('X(i,j)=X(%d,%d):%f, sum: %f: \n', i, j, X(i,j),sum(X(:,j)) )
        end
    end
    fprintf('p is: \n' )
    disp(p)
    
    %%???j??????e(j)
    k=1/log(n);
    %fprintf('k is: %f: \n', k)
    for j=1:m
        e(j)=-k*sum(p(:,j).*log(p(:,j))); 
        %fprintf('j is: %d.  p(:,j) is: \n', j ) 
        %disp( p(:,j) )
        %fprintf('log(p(:,j)) is: \n')
        %disp( log(p(:,j)) )
        %fprintf('p(:,j).*log(p(:,j)) is: \n')
        %disp( (p(:,j).*log(p(:,j))) )
        %fprintf('sum of them is: \n')
        %disp( sum(p(:,j).*log(p(:,j))) )
        %fprintf('e(j) is: %d \n', e(j) )
    end
    fprintf('e is:\n' )
    disp(e)
    
    
    d=ones(1,m)-e; %????????
    w=d./sum(d); 
    fprintf('=========================> w: \n')
    disp(w)
    
    s=100*w*X'; %?????
       
end