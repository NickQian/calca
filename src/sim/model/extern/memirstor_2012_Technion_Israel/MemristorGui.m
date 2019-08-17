%14.11.2011
%Model name: model of Memristor based on Theoretical formulas
%           this model was written by Dmitry Fliter and Keren Talisveyberg 
%           Technion Israel institute of technology EE faculty December 2011

function varargout = MemristorGui(varargin)
% MEMRISTORGUI MATLAB code for MemristorGui.fig
%      MEMRISTORGUI, by itself, creates a new MEMRISTORGUI or raises the existing
%      singleton*.
%
%      H = MEMRISTORGUI returns the handle to a new MEMRISTORGUI or the handle to
%      the existing singleton*.
%
%      MEMRISTORGUI('CALLBACK',hObject,eventData,handles,...) calls the local
%      function named CALLBACK in MEMRISTORGUI.M with the given input arguments.
%
%      MEMRISTORGUI('Property','Value',...) creates a new MEMRISTORGUI or raises the
%      existing singleton*.  Starting from the left, property value pairs are
%      applied to the GUI before MemristorGui_OpeningFcn gets called.  An
%      unrecognized property name or invalid value makes property application
%      stop.  All inputs are passed to MemristorGui_OpeningFcn via varargin.
%
%      *See GUI Options on GUIDE's Tools menu.  Choose "GUI allows only one
%      instance to run (singleton)".
%
% See also: GUIDE, GUIDATA, GUIHANDLES

% Edit the above text to modify the response to help MemristorGui

% Last Modified by GUIDE v2.5 18-Jan-2012 17:27:03

% Begin initialization code - DO NOT EDIT
gui_Singleton = 1;
gui_State = struct('gui_Name',       mfilename, ...
                   'gui_Singleton',  gui_Singleton, ...
                   'gui_OpeningFcn', @MemristorGui_OpeningFcn, ...
                   'gui_OutputFcn',  @MemristorGui_OutputFcn, ...
                   'gui_LayoutFcn',  [] , ...
                   'gui_Callback',   []);
if nargin && ischar(varargin{1})
    gui_State.gui_Callback = str2func(varargin{1});
end

if nargout
    [varargout{1:nargout}] = gui_mainfcn(gui_State, varargin{:});
else
    gui_mainfcn(gui_State, varargin{:});
end
% End initialization code - DO NOT EDIT


% --- Executes just before MemristorGui is made visible.
function MemristorGui_OpeningFcn(hObject, eventdata, handles, varargin)
% This function has no output args, see OutputFcn.
% hObject    handle to figure
% eventdata  reserved - to be defined in a future version of MATLAB
% handles    structure with handles and user data (see GUIDATA)
% varargin   command line arguments to MemristorGui (see VARARGIN)

% Choose default command line output for MemristorGui
handles.output = hObject;

%adds the standard toolbar
set(hObject,'toolbar','figure');

% Update handles structure
guidata(hObject, handles);

% UIWAIT makes MemristorGui wait for user response (see UIRESUME)
% uiwait(handles.figure1);


% --- Outputs from this function are returned to the command line.
function varargout = MemristorGui_OutputFcn(hObject, eventdata, handles) 
% varargout  cell array for returning output args (see VARARGOUT);
% hObject    handle to figure
% eventdata  reserved - to be defined in a future version of MATLAB
% handles    structure with handles and user data (see GUIDATA)

% Get default command line output from handles structure
varargout{1} = handles.output;

% --- Executes on selection change in model_popupmenu.
function model_popupmenu_Callback(hObject, eventdata, handles)
% hObject    handle to model_popupmenu (see GCBO)
% eventdata  reserved - to be defined in a future version of MATLAB
% handles    structure with handles and user data (see GUIDATA)

% Hints: contents = cellstr(get(hObject,'String')) returns model_popupmenu contents as cell array
%        contents{get(hObject,'Value')} returns selected item from model_popupmenu
%gets the selected option
switch get(handles.model_popupmenu,'Value')
    case 1
        model=0; %Linear Ion Drift
    case 2
        model=1; %Simmons Tunnel Barrier
    case 3
        model=2; %Team
    case 4
        model=3; % Nonlinear Ion Drift
    otherwise
end

% --- Executes during object creation, after setting all properties.
function model_popupmenu_CreateFcn(hObject, eventdata, handles)
% hObject    handle to model_popupmenu (see GCBO)
% eventdata  reserved - to be defined in a future version of MATLAB
% handles    empty - handles not created until after all CreateFcns called

% Hint: popupmenu controls usually have a white background on Windows.
%       See ISPC and COMPUTER.
if ispc && isequal(get(hObject,'BackgroundColor'), get(0,'defaultUicontrolBackgroundColor'))
    set(hObject,'BackgroundColor','white');
end

% --- Executes on selection change in window_popupmenu.
function window_popupmenu_Callback(hObject, eventdata, handles)
% hObject    handle to window_popupmenu (see GCBO)
% eventdata  reserved - to be defined in a future version of MATLAB
% handles    structure with handles and user data (see GUIDATA)

% Hints: contents = cellstr(get(hObject,'String')) returns window_popupmenu contents as cell array
%        contents{get(hObject,'Value')} returns selected item from window_popupmenu
%gets the selected option
switch get(handles.window_popupmenu,'Value')
    case 1
        win=0; %Ideal window
    case 2
        win=1; %Jogelkar window
    case 3
        win=2; %Biolek window
    case 4
        win=3; %Prodromakis window
    case 5
        win=4; %Kvatinsky window
    otherwise
end

% --- Executes during object creation, after setting all properties.
function window_popupmenu_CreateFcn(hObject, eventdata, handles)
% hObject    handle to window_popupmenu (see GCBO)
% eventdata  reserved - to be defined in a future version of MATLAB
% handles    empty - handles not created until after all CreateFcns called

% Hint: popupmenu controls usually have a white background on Windows.
%       See ISPC and COMPUTER.
if ispc && isequal(get(hObject,'BackgroundColor'), get(0,'defaultUicontrolBackgroundColor'))
    set(hObject,'BackgroundColor','white');
end

% --- Executes on selection change in iv_popupmenu.
function iv_popupmenu_Callback(hObject, eventdata, handles)
% hObject    handle to iv_popupmenu (see GCBO)
% eventdata  reserved - to be defined in a future version of MATLAB
% handles    structure with handles and user data (see GUIDATA)

% Hints: contents = cellstr(get(hObject,'String')) returns iv_popupmenu contents as cell array
%        contents{get(hObject,'Value')} returns selected item from iv_popupmenu
%gets the selected option
switch get(handles.iv_popupmenu,'Value')
    case 1
        iv=0; %V=IR
    case 2
        iv=1; %V=I*exp(...)
    otherwise
end

% --- Executes during object creation, after setting all properties.

function iv_popupmenu_CreateFcn(hObject, eventdata, handles)
% hObject    handle to iv_popupmenu (see GCBO)
% eventdata  reserved - to be defined in a future version of MATLAB
% handles    empty - handles not created until after all CreateFcns called

% Hint: popupmenu controls usually have a white background on Windows.
%       See ISPC and COMPUTER.
if ispc && isequal(get(hObject,'BackgroundColor'), get(0,'defaultUicontrolBackgroundColor'))
    set(hObject,'BackgroundColor','white');
end


%% global parameters

function frequency_editText_Callback(hObject, eventdata, handles)
% hObject    handle to frequency_editText (see GCBO)
% eventdata  reserved - to be defined in a future version of MATLAB
% handles    structure with handles and user data (see GUIDATA)

% Hints: get(hObject,'String') returns contents of frequency_editText as text
%        str2double(get(hObject,'String')) returns contents of frequency_editText as a double
%store the contents of frequency_editText as a string. if the string
%is not a number then input will be empty
freq = str2num(get(hObject,'String'));


%checks to see if input is empty. if so, default frequency_editText to zero
if (isempty(freq))
     set(hObject,'String','0')
end
guidata(hObject, handles);

% --- Executes during object creation, after setting all properties.
function frequency_editText_CreateFcn(hObject, eventdata, handles)
% hObject    handle to frequency_editText (see GCBO)
% eventdata  reserved - to be defined in a future version of MATLAB
% handles    empty - handles not created until after all CreateFcns called

% Hint: edit controls usually have a white background on Windows.
%       See ISPC and COMPUTER.
if ispc && isequal(get(hObject,'BackgroundColor'), get(0,'defaultUicontrolBackgroundColor'))
    set(hObject,'BackgroundColor','white');
end

function amplitude_editText_Callback(hObject, eventdata, handles)
% hObject    handle to amplitude_editText (see GCBO)
% eventdata  reserved - to be defined in a future version of MATLAB
% handles    structure with handles and user data (see GUIDATA)

% Hints: get(hObject,'String') returns contents of amplitude_editText as text
%        str2double(get(hObject,'String')) returns contents of amplitude_editText as a double
%store the contents of amplitude_editText as a string. if the string
%is not a number then input will be empty
amp = str2num(get(handles.amplitude_editText,'String'));
amp = str2num(get(hObject,'String'));
%checks to see if input is empty. if so, default amplitude_editText to zero
if (isempty(amp))
     set(hObject,'String','0')
end
guidata(hObject, handles);

% --- Executes during object creation, after setting all properties.
function amplitude_editText_CreateFcn(hObject, eventdata, handles)
% hObject    handle to amplitude_editText (see GCBO)
% eventdata  reserved - to be defined in a future version of MATLAB
% handles    empty - handles not created until after all CreateFcns called

% Hint: edit controls usually have a white background on Windows.
%       See ISPC and COMPUTER.
if ispc && isequal(get(hObject,'BackgroundColor'), get(0,'defaultUicontrolBackgroundColor'))
    set(hObject,'BackgroundColor','white');
end

function Ron_editText_Callback(hObject, eventdata, handles)
% hObject    handle to Ron_editText (see GCBO)
% eventdata  reserved - to be defined in a future version of MATLAB
% handles    structure with handles and user data (see GUIDATA)

% Hints: get(hObject,'String') returns contents of Ron_editText as text
%        str2double(get(hObject,'String')) returns contents of Ron_editText as a double
%store the contents of Ron_editText as a string. if the string
%is not a number then input will be empty
Ron = str2num(get(hObject,'String'));

%checks to see if input is empty. if so, default Ron_editText to zero
if (isempty(Ron))
     set(hObject,'String','0')
end
guidata(hObject, handles);

% --- Executes during object creation, after setting all properties.
function Ron_editText_CreateFcn(hObject, eventdata, handles)
% hObject    handle to Ron_editText (see GCBO)
% eventdata  reserved - to be defined in a future version of MATLAB
% handles    empty - handles not created until after all CreateFcns called

% Hint: edit controls usually have a white background on Windows.
%       See ISPC and COMPUTER.
if ispc && isequal(get(hObject,'BackgroundColor'), get(0,'defaultUicontrolBackgroundColor'))
    set(hObject,'BackgroundColor','white');
end

function Roff_editText_Callback(hObject, eventdata, handles)
% hObject    handle to Roff_editText (see GCBO)
% eventdata  reserved - to be defined in a future version of MATLAB
% handles    structure with handles and user data (see GUIDATA)

% Hints: get(hObject,'String') returns contents of Roff_editText as text
%        str2double(get(hObject,'String')) returns contents of Roff_editText as a double
%store the contents of Roff_editText as a string. if the string
%is not a number then input will be empty
Roff = str2num(get(hObject,'String'));

%checks to see if input is empty. if so, default Roff_editText to zero
if (isempty(Roff))
     set(hObject,'String','0')
end
guidata(hObject, handles);


% --- Executes during object creation, after setting all properties.
function Roff_editText_CreateFcn(hObject, eventdata, handles)
% hObject    handle to Roff_editText (see GCBO)
% eventdata  reserved - to be defined in a future version of MATLAB
% handles    empty - handles not created until after all CreateFcns called

% Hint: edit controls usually have a white background on Windows.
%       See ISPC and COMPUTER.
if ispc && isequal(get(hObject,'BackgroundColor'), get(0,'defaultUicontrolBackgroundColor'))
    set(hObject,'BackgroundColor','white');
end

function D_editText_Callback(hObject, eventdata, handles)
% hObject    handle to D_editText (see GCBO)
% eventdata  reserved - to be defined in a future version of MATLAB
% handles    structure with handles and user data (see GUIDATA)

% Hints: get(hObject,'String') returns contents of D_editText as text
%        str2double(get(hObject,'String')) returns contents of D_editText as a double
%store the contents of D_editText as a string. if the string
%is not a number then input will be empty
D = str2num(get(hObject,'String'));

%checks to see if input is empty. if so, default D_editText to zero
if (isempty(D))
     set(hObject,'String','0')
end
guidata(hObject, handles);


% --- Executes during object creation, after setting all properties.
function D_editText_CreateFcn(hObject, eventdata, handles)
% hObject    handle to D_editText (see GCBO)
% eventdata  reserved - to be defined in a future version of MATLAB
% handles    empty - handles not created until after all CreateFcns called

% Hint: edit controls usually have a white background on Windows.
%       See ISPC and COMPUTER.
if ispc && isequal(get(hObject,'BackgroundColor'), get(0,'defaultUicontrolBackgroundColor'))
    set(hObject,'BackgroundColor','white');
end


function uV_editText_Callback(hObject, eventdata, handles)
% hObject    handle to uV_editText (see GCBO)
% eventdata  reserved - to be defined in a future version of MATLAB
% handles    structure with handles and user data (see GUIDATA)

% Hints: get(hObject,'String') returns contents of uV_editText as text
%        str2double(get(hObject,'String')) returns contents of uV_editText as a double
%store the contents of uV_editText as a string. if the string
%is not a number then input will be empty
uV = str2num(get(hObject,'String'));

%checks to see if input is empty. if so, default uV_editText to zero
if (isempty(uV))
     set(hObject,'String','0')
end
guidata(hObject, handles);

% --- Executes during object creation, after setting all properties.
function uV_editText_CreateFcn(hObject, eventdata, handles)
% hObject    handle to uV_editText (see GCBO)
% eventdata  reserved - to be defined in a future version of MATLAB
% handles    empty - handles not created until after all CreateFcns called

% Hint: edit controls usually have a white background on Windows.
%       See ISPC and COMPUTER.
if ispc && isequal(get(hObject,'BackgroundColor'), get(0,'defaultUicontrolBackgroundColor'))
    set(hObject,'BackgroundColor','white');
end

function V_t_editText_Callback(hObject, eventdata, handles)
% hObject    handle to V_t_editText (see GCBO)
% eventdata  reserved - to be defined in a future version of MATLAB
% handles    structure with handles and user data (see GUIDATA)

% Hints: get(hObject,'String') returns contents of V_t_editText as text
%        str2double(get(hObject,'String')) returns contents of V_t_editText as a double
%store the contents of V_t_editText as a string. if the string
%is not a number then input will be empty
V_t = str2num(get(hObject,'String'));

%checks to see if input is empty. if so, default V_t_editText to zero
if (isempty(V_t))
     set(hObject,'String','0')
end
guidata(hObject, handles);


% --- Executes during object creation, after setting all properties.
function V_t_editText_CreateFcn(hObject, eventdata, handles)
% hObject    handle to V_t_editText (see GCBO)
% eventdata  reserved - to be defined in a future version of MATLAB
% handles    empty - handles not created until after all CreateFcns called

% Hint: edit controls usually have a white background on Windows.
%       See ISPC and COMPUTER.
if ispc && isequal(get(hObject,'BackgroundColor'), get(0,'defaultUicontrolBackgroundColor'))
    set(hObject,'BackgroundColor','white');
end


function p_coeff_editText_Callback(hObject, eventdata, handles)
% hObject    handle to p_coeff_editText (see GCBO)
% eventdata  reserved - to be defined in a future version of MATLAB
% handles    structure with handles and user data (see GUIDATA)

% Hints: get(hObject,'String') returns contents of p_coeff_editText as text
%        str2double(get(hObject,'String')) returns contents of p_coeff_editText as a double
%store the contents of p_coeff_editText as a string. if the string
%is not a number then input will be empty
P_coeff = str2num(get(hObject,'String'));

%checks to see if input is empty. if so, default p_coeff_editText to zero
if (isempty(P_coeff))
     set(hObject,'String','0')
end
guidata(hObject, handles);

% --- Executes during object creation, after setting all properties.
function p_coeff_editText_CreateFcn(hObject, eventdata, handles)
% hObject    handle to p_coeff_editText (see GCBO)
% eventdata  reserved - to be defined in a future version of MATLAB
% handles    empty - handles not created until after all CreateFcns called

% Hint: edit controls usually have a white background on Windows.
%       See ISPC and COMPUTER.
if ispc && isequal(get(hObject,'BackgroundColor'), get(0,'defaultUicontrolBackgroundColor'))
    set(hObject,'BackgroundColor','white');
end



function W0_editText_Callback(hObject, eventdata, handles)
% hObject    handle to W0_editText (see GCBO)
% eventdata  reserved - to be defined in a future version of MATLAB
% handles    structure with handles and user data (see GUIDATA)

% Hints: get(hObject,'String') returns contents of W0_editText as text
%        str2double(get(hObject,'String')) returns contents of W0_editText as a double
%store the contents of W0_editText as a string. if the string
%is not a number then input will be empty
w_init = str2num(get(hObject,'String'));

%checks to see if input is empty. if so, default W0_editText to zero
if (isempty(w_init))
     set(hObject,'String','0')
end
guidata(hObject, handles);

% --- Executes during object creation, after setting all properties.
function W0_editText_CreateFcn(hObject, eventdata, handles)
% hObject    handle to W0_editText (see GCBO)
% eventdata  reserved - to be defined in a future version of MATLAB
% handles    empty - handles not created until after all CreateFcns called

% Hint: edit controls usually have a white background on Windows.
%       See ISPC and COMPUTER.
if ispc && isequal(get(hObject,'BackgroundColor'), get(0,'defaultUicontrolBackgroundColor'))
    set(hObject,'BackgroundColor','white');
end


function J_editText_Callback(hObject, eventdata, handles)
% hObject    handle to J_editText (see GCBO)
% eventdata  reserved - to be defined in a future version of MATLAB
% handles    structure with handles and user data (see GUIDATA)

% Hints: get(hObject,'String') returns contents of J_editText as text
%        str2double(get(hObject,'String')) returns contents of J_editText as a double
%store the contents of J_editText as a string. if the string
%is not a number then input will be empty
J = str2num(get(hObject,'String'));

%checks to see if input is empty. if so, default J_editText to zero
if (isempty(J))
     set(hObject,'String','0')
end
guidata(hObject, handles);

% --- Executes during object creation, after setting all properties.
function J_editText_CreateFcn(hObject, eventdata, handles)
% hObject    handle to J_editText (see GCBO)
% eventdata  reserved - to be defined in a future version of MATLAB
% handles    empty - handles not created until after all CreateFcns called

% Hint: edit controls usually have a white background on Windows.
%       See ISPC and COMPUTER.
if ispc && isequal(get(hObject,'BackgroundColor'), get(0,'defaultUicontrolBackgroundColor'))
    set(hObject,'BackgroundColor','white');
end

function num_of_cycles_editText_Callback(hObject, eventdata, handles)
% hObject    handle to num_of_cycles_editText (see GCBO)
% eventdata  reserved - to be defined in a future version of MATLAB
% handles    structure with handles and user data (see GUIDATA)

% Hints: get(hObject,'String') returns contents of num_of_cycles_editText as text
%        str2double(get(hObject,'String')) returns contents of num_of_cycles_editText as a double
%store the contents of num_of_cycles_editText as a string. if the string
%is not a number then input will be empty
num_of_cycles = str2num(get(hObject,'String'));

%checks to see if input is empty. if so, default num_of_cycles_editText to zero
if (isempty(num_of_cycles))
     set(hObject,'String','0')
end
guidata(hObject, handles);


% --- Executes during object creation, after setting all properties.
function num_of_cycles_editText_CreateFcn(hObject, eventdata, handles)
% hObject    handle to num_of_cycles_editText (see GCBO)
% eventdata  reserved - to be defined in a future version of MATLAB
% handles    empty - handles not created until after all CreateFcns called

% Hint: edit controls usually have a white background on Windows.
%       See ISPC and COMPUTER.
if ispc && isequal(get(hObject,'BackgroundColor'), get(0,'defaultUicontrolBackgroundColor'))
    set(hObject,'BackgroundColor','white');
end

%% Simmons Tunnel Barrier & Team Parameters


function a_on_editText_Callback(hObject, eventdata, handles)
% hObject    handle to a_on_editText (see GCBO)
% eventdata  reserved - to be defined in a future version of MATLAB
% handles    structure with handles and user data (see GUIDATA)

% Hints: get(hObject,'String') returns contents of a_on_editText as text
%        str2double(get(hObject,'String')) returns contents of a_on_editText as a double
%store the contents of a_on_editText as a string. if the string
%is not a number then input will be empty
 a_on = str2num(get(hObject,'String'));

%checks to see if input is empty. if so, default a_on_editText to zero
if (isempty(a_on))
     set(hObject,'String','0')
end
guidata(hObject, handles);

% --- Executes during object creation, after setting all properties.
function a_on_editText_CreateFcn(hObject, eventdata, handles)
% hObject    handle to a_on_editText (see GCBO)
% eventdata  reserved - to be defined in a future version of MATLAB
% handles    empty - handles not created until after all CreateFcns called

% Hint: edit controls usually have a white background on Windows.
%       See ISPC and COMPUTER.
if ispc && isequal(get(hObject,'BackgroundColor'), get(0,'defaultUicontrolBackgroundColor'))
    set(hObject,'BackgroundColor','white');
end

function a_off_editText_Callback(hObject, eventdata, handles)
% hObject    handle to a_off_editText (see GCBO)
% eventdata  reserved - to be defined in a future version of MATLAB
% handles    structure with handles and user data (see GUIDATA)

% Hints: get(hObject,'String') returns contents of a_off_editText as text
%        str2double(get(hObject,'String')) returns contents of a_off_editText as a double
%store the contents of a_off_editText as a string. if the string
%is not a number then input will be empty
 a_off = str2num(get(hObject,'String'));

%checks to see if input is empty. if so, default a_off_editText to zero
if (isempty(a_off))
     set(hObject,'String','0')
end
guidata(hObject, handles);

% --- Executes during object creation, after setting all properties.
function a_off_editText_CreateFcn(hObject, eventdata, handles)
% hObject    handle to a_off_editText (see GCBO)
% eventdata  reserved - to be defined in a future version of MATLAB
% handles    empty - handles not created until after all CreateFcns called

% Hint: edit controls usually have a white background on Windows.
%       See ISPC and COMPUTER.
if ispc && isequal(get(hObject,'BackgroundColor'), get(0,'defaultUicontrolBackgroundColor'))
    set(hObject,'BackgroundColor','white');
end

function i_on_editText_Callback(hObject, eventdata, handles)
% hObject    handle to i_on_editText (see GCBO)
% eventdata  reserved - to be defined in a future version of MATLAB
% handles    structure with handles and user data (see GUIDATA)

% Hints: get(hObject,'String') returns contents of i_on_editText as text
%        str2double(get(hObject,'String')) returns contents of i_on_editText as a double
%store the contents of i_on_editText as a string. if the string
%is not a number then input will be empty
 i_on = str2num(get(hObject,'String'));

%checks to see if input is empty. if so, default i_on_editText to zero
if (isempty(i_on))
     set(hObject,'String','0')
end
guidata(hObject, handles);

% --- Executes during object creation, after setting all properties.
function i_on_editText_CreateFcn(hObject, eventdata, handles)
% hObject    handle to i_on_editText (see GCBO)
% eventdata  reserved - to be defined in a future version of MATLAB
% handles    empty - handles not created until after all CreateFcns called

% Hint: edit controls usually have a white background on Windows.
%       See ISPC and COMPUTER.
if ispc && isequal(get(hObject,'BackgroundColor'), get(0,'defaultUicontrolBackgroundColor'))
    set(hObject,'BackgroundColor','white');
end

function i_off_editText_Callback(hObject, eventdata, handles)
% hObject    handle to i_off_editText (see GCBO)
% eventdata  reserved - to be defined in a future version of MATLAB
% handles    structure with handles and user data (see GUIDATA)

% Hints: get(hObject,'String') returns contents of i_off_editText as text
%        str2double(get(hObject,'String')) returns contents of i_off_editText as a double
%store the contents of i_off_editText as a string. if the string
%is not a number then input will be empty
 i_off = str2num(get(hObject,'String'));

%checks to see if input is empty. if so, default i_off_editText to zero
if (isempty(i_off))
     set(hObject,'String','0')
end
guidata(hObject, handles);

% --- Executes during object creation, after setting all properties.
function i_off_editText_CreateFcn(hObject, eventdata, handles)
% hObject    handle to i_off_editText (see GCBO)
% eventdata  reserved - to be defined in a future version of MATLAB
% handles    empty - handles not created until after all CreateFcns called

% Hint: edit controls usually have a white background on Windows.
%       See ISPC and COMPUTER.
if ispc && isequal(get(hObject,'BackgroundColor'), get(0,'defaultUicontrolBackgroundColor'))
    set(hObject,'BackgroundColor','white');
end

function c_on_editText_Callback(hObject, eventdata, handles)
% hObject    handle to c_on_editText (see GCBO)
% eventdata  reserved - to be defined in a future version of MATLAB
% handles    structure with handles and user data (see GUIDATA)

% Hints: get(hObject,'String') returns contents of c_on_editText as text
%        str2double(get(hObject,'String')) returns contents of c_on_editText as a double
%store the contents of c_on_editText as a string. if the string
%is not a number then input will be empty
 c_on = str2num(get(hObject,'String'));

%checks to see if input is empty. if so, default c_on_editText to zero
if (isempty(c_on))
     set(hObject,'String','0')
end
guidata(hObject, handles);

% --- Executes during object creation, after setting all properties.
function c_on_editText_CreateFcn(hObject, eventdata, handles)
% hObject    handle to c_on_editText (see GCBO)
% eventdata  reserved - to be defined in a future version of MATLAB
% handles    empty - handles not created until after all CreateFcns called

% Hint: edit controls usually have a white background on Windows.
%       See ISPC and COMPUTER.
if ispc && isequal(get(hObject,'BackgroundColor'), get(0,'defaultUicontrolBackgroundColor'))
    set(hObject,'BackgroundColor','white');
end

function c_off_editText_Callback(hObject, eventdata, handles)
% hObject    handle to c_off_editText (see GCBO)
% eventdata  reserved - to be defined in a future version of MATLAB
% handles    structure with handles and user data (see GUIDATA)

% Hints: get(hObject,'String') returns contents of c_off_editText as text
%        str2double(get(hObject,'String')) returns contents of c_off_editText as a double
%store the contents of c_off_editText as a string. if the string
%is not a number then input will be empty
c_off = str2num(get(hObject,'String'));

%checks to see if input is empty. if so, default c_off_editText to zero
if (isempty(c_off))
     set(hObject,'String','0')
end
guidata(hObject, handles);


% --- Executes during object creation, after setting all properties.
function c_off_editText_CreateFcn(hObject, eventdata, handles)
% hObject    handle to c_off_editText (see GCBO)
% eventdata  reserved - to be defined in a future version of MATLAB
% handles    empty - handles not created until after all CreateFcns called

% Hint: edit controls usually have a white background on Windows.
%       See ISPC and COMPUTER.
if ispc && isequal(get(hObject,'BackgroundColor'), get(0,'defaultUicontrolBackgroundColor'))
    set(hObject,'BackgroundColor','white');
end



function b_editText_Callback(hObject, eventdata, handles)
% hObject    handle to b_editText (see GCBO)
% eventdata  reserved - to be defined in a future version of MATLAB
% handles    structure with handles and user data (see GUIDATA)

% Hints: get(hObject,'String') returns contents of b_editText as text
%        str2double(get(hObject,'String')) returns contents of b_editText as a double
%store the contents of b_editText as a string. if the string
%is not a number then input will be empty
b = str2num(get(hObject,'String'));

%checks to see if input is empty. if so, default b_editText to zero
if (isempty(b))
     set(hObject,'String','0')
end
guidata(hObject, handles);

% --- Executes during object creation, after setting all properties.
function b_editText_CreateFcn(hObject, eventdata, handles)
% hObject    handle to b_editText (see GCBO)
% eventdata  reserved - to be defined in a future version of MATLAB
% handles    empty - handles not created until after all CreateFcns called

% Hint: edit controls usually have a white background on Windows.
%       See ISPC and COMPUTER.
if ispc && isequal(get(hObject,'BackgroundColor'), get(0,'defaultUicontrolBackgroundColor'))
    set(hObject,'BackgroundColor','white');
end


function X_c_editText_Callback(hObject, eventdata, handles)
% hObject    handle to X_c_editText (see GCBO)
% eventdata  reserved - to be defined in a future version of MATLAB
% handles    structure with handles and user data (see GUIDATA)

% Hints: get(hObject,'String') returns contents of X_c_editText as text
%        str2double(get(hObject,'String')) returns contents of X_c_editText as a double
%store the contents of X_c_editText as a string. if the string
%is not a number then input will be empty
 X_c = str2num(get(hObject,'String'));

%checks to see if input is empty. if so, default X_c_editText to zero
if (isempty(X_c))
     set(hObject,'String','0')
end
guidata(hObject, handles);

% --- Executes during object creation, after setting all properties.
function X_c_editText_CreateFcn(hObject, eventdata, handles)
% hObject    handle to X_c_editText (see GCBO)
% eventdata  reserved - to be defined in a future version of MATLAB
% handles    empty - handles not created until after all CreateFcns called

% Hint: edit controls usually have a white background on Windows.
%       See ISPC and COMPUTER.
if ispc && isequal(get(hObject,'BackgroundColor'), get(0,'defaultUicontrolBackgroundColor'))
    set(hObject,'BackgroundColor','white');
end

function k_on_editText_Callback(hObject, eventdata, handles)
% hObject    handle to k_on_editText (see GCBO)
% eventdata  reserved - to be defined in a future version of MATLAB
% handles    structure with handles and user data (see GUIDATA)

% Hints: get(hObject,'String') returns contents of k_on_editText as text
%        str2double(get(hObject,'String')) returns contents of k_on_editText as a double
%store the contents of k_on_editText as a string. if the string
%is not a number then input will be empty
k_on = str2num(get(hObject,'String'));

%checks to see if input is empty. if so, default k_on_editText to zero
if (isempty(k_on))
     set(hObject,'String','0')
end
guidata(hObject, handles);

% --- Executes during object creation, after setting all properties.
function k_on_editText_CreateFcn(hObject, eventdata, handles)
% hObject    handle to k_on_editText (see GCBO)
% eventdata  reserved - to be defined in a future version of MATLAB
% handles    empty - handles not created until after all CreateFcns called

% Hint: edit controls usually have a white background on Windows.
%       See ISPC and COMPUTER.
if ispc && isequal(get(hObject,'BackgroundColor'), get(0,'defaultUicontrolBackgroundColor'))
    set(hObject,'BackgroundColor','white');
end

function k_off_editText_Callback(hObject, eventdata, handles)
% hObject    handle to k_off_editText (see GCBO)
% eventdata  reserved - to be defined in a future version of MATLAB
% handles    structure with handles and user data (see GUIDATA)

% Hints: get(hObject,'String') returns contents of k_off_editText as text
%        str2double(get(hObject,'String')) returns contents of k_off_editText as a double
%store the contents of k_off_editText as a string. if the string
%is not a number then input will be empty
k_off = str2num(get(hObject,'String'));

%checks to see if input is empty. if so, default k_off_editText to zero
if (isempty(k_off))
     set(hObject,'String','0')
end
guidata(hObject, handles);

% --- Executes during object creation, after setting all properties.
function k_off_editText_CreateFcn(hObject, eventdata, handles)
% hObject    handle to k_off_editText (see GCBO)
% eventdata  reserved - to be defined in a future version of MATLAB
% handles    empty - handles not created until after all CreateFcns called

% Hint: edit controls usually have a white background on Windows.
%       See ISPC and COMPUTER.
if ispc && isequal(get(hObject,'BackgroundColor'), get(0,'defaultUicontrolBackgroundColor'))
    set(hObject,'BackgroundColor','white');
end

function alpha_on_editText_Callback(hObject, eventdata, handles)
% hObject    handle to alpha_on_editText (see GCBO)
% eventdata  reserved - to be defined in a future version of MATLAB
% handles    structure with handles and user data (see GUIDATA)

% Hints: get(hObject,'String') returns contents of alpha_on_editText as text
%        str2double(get(hObject,'String')) returns contents of alpha_on_editText as a double
%store the contents of alpha_on_editText as a string. if the string
%is not a number then input will be empty
alpha_on = str2num(get(hObject,'String'));

%checks to see if input is empty. if so, default alpha_on_editText to zero
if (isempty(alpha_on))
     set(hObject,'String','0')
end
guidata(hObject, handles);

% --- Executes during object creation, after setting all properties.
function alpha_on_editText_CreateFcn(hObject, eventdata, handles)
% hObject    handle to alpha_on_editText (see GCBO)
% eventdata  reserved - to be defined in a future version of MATLAB
% handles    empty - handles not created until after all CreateFcns called

% Hint: edit controls usually have a white background on Windows.
%       See ISPC and COMPUTER.
if ispc && isequal(get(hObject,'BackgroundColor'), get(0,'defaultUicontrolBackgroundColor'))
    set(hObject,'BackgroundColor','white');
end

function alpha_off_editText_Callback(hObject, eventdata, handles)
% hObject    handle to alpha_off_editText (see GCBO)
% eventdata  reserved - to be defined in a future version of MATLAB
% handles    structure with handles and user data (see GUIDATA)

% Hints: get(hObject,'String') returns contents of alpha_off_editText as text
%        str2double(get(hObject,'String')) returns contents of alpha_off_editText as a double
%store the contents of alpha_off_editText as a string. if the string
%is not a number then input will be empty
alpha_off = str2num(get(hObject,'String'));

%checks to see if input is empty. if so, default alpha_off_editText to zero
if (isempty(alpha_off))
     set(hObject,'String','0')
end
guidata(hObject, handles);

% --- Executes during object creation, after setting all properties.
function alpha_off_editText_CreateFcn(hObject, eventdata, handles)
% hObject    handle to alpha_off_editText (see GCBO)
% eventdata  reserved - to be defined in a future version of MATLAB
% handles    empty - handles not created until after all CreateFcns called

% Hint: edit controls usually have a white background on Windows.
%       See ISPC and COMPUTER.
if ispc && isequal(get(hObject,'BackgroundColor'), get(0,'defaultUicontrolBackgroundColor'))
    set(hObject,'BackgroundColor','white');
end


function x_on_editText_Callback(hObject, eventdata, handles)
% hObject    handle to x_on_editText (see GCBO)
% eventdata  reserved - to be defined in a future version of MATLAB
% handles    structure with handles and user data (see GUIDATA)

% Hints: get(hObject,'String') returns contents of x_on_editText as text
%        str2double(get(hObject,'String')) returns contents of x_on_editText as a double
x_on = str2num(get(hObject,'String'));

%checks to see if input is empty. if so, default alpha_off_editText to zero
if (isempty(x_on))
     set(hObject,'String','0')
end
guidata(hObject, handles);

% --- Executes during object creation, after setting all properties.
function x_on_editText_CreateFcn(hObject, eventdata, handles)
% hObject    handle to x_on_editText (see GCBO)
% eventdata  reserved - to be defined in a future version of MATLAB
% handles    empty - handles not created until after all CreateFcns called

% Hint: edit controls usually have a white background on Windows.
%       See ISPC and COMPUTER.
if ispc && isequal(get(hObject,'BackgroundColor'), get(0,'defaultUicontrolBackgroundColor'))
    set(hObject,'BackgroundColor','white');
end



function x_off_editText_Callback(hObject, eventdata, handles)
% hObject    handle to x_off_editText (see GCBO)
% eventdata  reserved - to be defined in a future version of MATLAB
% handles    structure with handles and user data (see GUIDATA)

% Hints: get(hObject,'String') returns contents of x_off_editText as text
%        str2double(get(hObject,'String')) returns contents of x_off_editText as a double
x_off = str2num(get(hObject,'String'));

%checks to see if input is empty. if so, default alpha_off_editText to zero
if (isempty(x_off))
     set(hObject,'String','0')
end
guidata(hObject, handles);

% --- Executes during object creation, after setting all properties.
function x_off_editText_CreateFcn(hObject, eventdata, handles)
% hObject    handle to x_off_editText (see GCBO)
% eventdata  reserved - to be defined in a future version of MATLAB
% handles    empty - handles not created until after all CreateFcns called

% Hint: edit controls usually have a white background on Windows.
%       See ISPC and COMPUTER.
if ispc && isequal(get(hObject,'BackgroundColor'), get(0,'defaultUicontrolBackgroundColor'))
    set(hObject,'BackgroundColor','white');
end

%% Nonlinear Ion Drift parameters


function alpha_editText_Callback(hObject, eventdata, handles)
% hObject    handle to alpha_editText (see GCBO)
% eventdata  reserved - to be defined in a future version of MATLAB
% handles    structure with handles and user data (see GUIDATA)

% Hints: get(hObject,'String') returns contents of alpha_editText as text
%        str2double(get(hObject,'String')) returns contents of alpha_editText as a double
%store the contents of alpha_editText as a string. if the string
%is not a number then input will be empty
alpha = str2num(get(hObject,'String'));

%checks to see if input is empty. if so, default alpha_editText to zero
if (isempty(alpha))
     set(hObject,'String','0')
end
guidata(hObject, handles);

% --- Executes during object creation, after setting all properties.
function alpha_editText_CreateFcn(hObject, eventdata, handles)
% hObject    handle to alpha_editText (see GCBO)
% eventdata  reserved - to be defined in a future version of MATLAB
% handles    empty - handles not created until after all CreateFcns called

% Hint: edit controls usually have a white background on Windows.
%       See ISPC and COMPUTER.
if ispc && isequal(get(hObject,'BackgroundColor'), get(0,'defaultUicontrolBackgroundColor'))
    set(hObject,'BackgroundColor','white');
end

function beta_editText_Callback(hObject, eventdata, handles)
% hObject    handle to beta_editText (see GCBO)
% eventdata  reserved - to be defined in a future version of MATLAB
% handles    structure with handles and user data (see GUIDATA)

% Hints: get(hObject,'String') returns contents of beta_editText as text
%        str2double(get(hObject,'String')) returns contents of beta_editText as a double
%store the contents of beta_editText as a string. if the string
%is not a number then input will be empty
beta = str2num(get(hObject,'String'));

%checks to see if input is empty. if so, default beta_editText to zero
if (isempty(beta))
     set(hObject,'String','0')
end
guidata(hObject, handles);


% --- Executes during object creation, after setting all properties.
function beta_editText_CreateFcn(hObject, eventdata, handles)
% hObject    handle to beta_editText (see GCBO)
% eventdata  reserved - to be defined in a future version of MATLAB
% handles    empty - handles not created until after all CreateFcns called

% Hint: edit controls usually have a white background on Windows.
%       See ISPC and COMPUTER.
if ispc && isequal(get(hObject,'BackgroundColor'), get(0,'defaultUicontrolBackgroundColor'))
    set(hObject,'BackgroundColor','white');
end

function c_editText_Callback(hObject, eventdata, handles)
% hObject    handle to c_editText (see GCBO)
% eventdata  reserved - to be defined in a future version of MATLAB
% handles    structure with handles and user data (see GUIDATA)

% Hints: get(hObject,'String') returns contents of c_editText as text
%        str2double(get(hObject,'String')) returns contents of c_editText as a double
%store the contents of c_editText as a string. if the string
%is not a number then input will be empty
c = str2num(get(hObject,'String'));

%checks to see if input is empty. if so, default c_editText to zero
if (isempty(c))
     set(hObject,'String','0')
end
guidata(hObject, handles);

% --- Executes during object creation, after setting all properties.
function c_editText_CreateFcn(hObject, eventdata, handles)
% hObject    handle to c_editText (see GCBO)
% eventdata  reserved - to be defined in a future version of MATLAB
% handles    empty - handles not created until after all CreateFcns called

% Hint: edit controls usually have a white background on Windows.
%       See ISPC and COMPUTER.
if ispc && isequal(get(hObject,'BackgroundColor'), get(0,'defaultUicontrolBackgroundColor'))
    set(hObject,'BackgroundColor','white');
end


function g_editText_Callback(hObject, eventdata, handles)
% hObject    handle to g_editText (see GCBO)
% eventdata  reserved - to be defined in a future version of MATLAB
% handles    structure with handles and user data (see GUIDATA)

% Hints: get(hObject,'String') returns contents of g_editText as text
%        str2double(get(hObject,'String')) returns contents of g_editText as a double
%store the contents of g_editText as a string. if the string
%is not a number then input will be empty
g = str2num(get(hObject,'String'));

%checks to see if input is empty. if so, default g_editText to zero
if (isempty(g))
     set(hObject,'String','0')
end
guidata(hObject, handles);

% --- Executes during object creation, after setting all properties.
function g_editText_CreateFcn(hObject, eventdata, handles)
% hObject    handle to g_editText (see GCBO)
% eventdata  reserved - to be defined in a future version of MATLAB
% handles    empty - handles not created until after all CreateFcns called

% Hint: edit controls usually have a white background on Windows.
%       See ISPC and COMPUTER.
if ispc && isequal(get(hObject,'BackgroundColor'), get(0,'defaultUicontrolBackgroundColor'))
    set(hObject,'BackgroundColor','white');
end

function n_editText_Callback(hObject, eventdata, handles)
% hObject    handle to n_editText (see GCBO)
% eventdata  reserved - to be defined in a future version of MATLAB
% handles    structure with handles and user data (see GUIDATA)

% Hints: get(hObject,'String') returns contents of n_editText as text
%        str2double(get(hObject,'String')) returns contents of n_editText as a double
%store the contents of n_editText as a string. if the string
%is not a number then input will be empty
n = str2num(get(hObject,'String'));

%checks to see if input is empty. if so, default n_editText to zero
if (isempty(n))
     set(hObject,'String','0')
end
guidata(hObject, handles);

% --- Executes during object creation, after setting all properties.
function n_editText_CreateFcn(hObject, eventdata, handles)
% hObject    handle to n_editText (see GCBO)
% eventdata  reserved - to be defined in a future version of MATLAB
% handles    empty - handles not created until after all CreateFcns called

% Hint: edit controls usually have a white background on Windows.
%       See ISPC and COMPUTER.
if ispc && isequal(get(hObject,'BackgroundColor'), get(0,'defaultUicontrolBackgroundColor'))
    set(hObject,'BackgroundColor','white');
end

function q_editText_Callback(hObject, eventdata, handles)
% hObject    handle to q_editText (see GCBO)
% eventdata  reserved - to be defined in a future version of MATLAB
% handles    structure with handles and user data (see GUIDATA)

% Hints: get(hObject,'String') returns contents of q_editText as text
%        str2double(get(hObject,'String')) returns contents of q_editText as a double
%store the contents of q_editText as a string. if the string
%is not a number then input will be empty
q = str2num(get(hObject,'String'));

%checks to see if input is empty. if so, default q_editText to zero
if (isempty(q))
     set(hObject,'String','0')
end
guidata(hObject, handles);

% --- Executes during object creation, after setting all properties.
function q_editText_CreateFcn(hObject, eventdata, handles)
% hObject    handle to q_editText (see GCBO)
% eventdata  reserved - to be defined in a future version of MATLAB
% handles    empty - handles not created until after all CreateFcns called

% Hint: edit controls usually have a white background on Windows.
%       See ISPC and COMPUTER.
if ispc && isequal(get(hObject,'BackgroundColor'), get(0,'defaultUicontrolBackgroundColor'))
    set(hObject,'BackgroundColor','white');
end

function a_editText_Callback(hObject, eventdata, handles)
% hObject    handle to a_editText (see GCBO)
% eventdata  reserved - to be defined in a future version of MATLAB
% handles    structure with handles and user data (see GUIDATA)

% Hints: get(hObject,'String') returns contents of a_editText as text
%        str2double(get(hObject,'String')) returns contents of a_editText as a double
%store the contents of a_editText as a string. if the string
%is not a number then input will be empty
a = str2num(get(hObject,'String'));

%checks to see if input is empty. if so, default a_editText to zero
if (isempty(a))
     set(hObject,'String','0')
end
guidata(hObject, handles);

% --- Executes during object creation, after setting all properties.
function a_editText_CreateFcn(hObject, eventdata, handles)
% hObject    handle to a_editText (see GCBO)
% eventdata  reserved - to be defined in a future version of MATLAB
% handles    empty - handles not created until after all CreateFcns called

% Hint: edit controls usually have a white background on Windows.
%       See ISPC and COMPUTER.
if ispc && isequal(get(hObject,'BackgroundColor'), get(0,'defaultUicontrolBackgroundColor'))
    set(hObject,'BackgroundColor','white');
end


%%
% --- Executes on button press in plotW_pushbutton.
function plotW_pushbutton_Callback(hObject, eventdata, handles)
% hObject    handle to plotW_pushbutton (see GCBO)
% eventdata  reserved - to be defined in a future version of MATLAB
% handles    structure with handles and user data (see GUIDATA)

%selects axes1 as the current axes, so that
%Matlab knows where to plot the data

axes(handles.axes1)


amp = str2num(get(handles.amplitude_editText,'String'));
Roff = str2num(get(handles.Roff_editText,'String'));
Ron = str2num(get(handles.Ron_editText,'String'));
freq = str2num(get(handles.frequency_editText,'String'));
D = str2num(get(handles.D_editText,'String'));
uV =  str2num(get(handles.uV_editText,'String'));
V_t = str2num(get(handles.V_t_editText,'String'));
P_coeff = str2num(get(handles.p_coeff_editText,'String'));
w_init = str2num(get(handles.W0_editText,'String'));
J = str2num(get(handles.J_editText,'String'));
a_on = str2num(get(handles.a_on_editText,'String'));
a_off = str2num(get(handles.a_off_editText,'String'));
c_on = str2num(get(handles.c_on_editText,'String'));
c_off = str2num(get(handles.c_off_editText,'String'));
alpha_on = str2num(get(handles.alpha_on_editText,'String'));
alpha_off = str2num(get(handles.alpha_off_editText,'String'));
k_on = str2num(get(handles.k_on_editText,'String'));
k_off = str2num(get(handles.k_off_editText,'String'));
i_on = str2num(get(handles.i_on_editText,'String'));
i_off = str2num(get(handles.i_off_editText,'String'));
x_on = str2num(get(handles.x_on_editText,'String'));
x_off = str2num(get(handles.x_off_editText,'String'));
beta = str2num(get(handles.beta_editText,'String'));
a = str2num(get(handles.a_editText,'String'));
c = str2num(get(handles.c_editText,'String'));
n = str2num(get(handles.n_editText,'String'));
q = str2num(get(handles.q_editText,'String'));
g = str2num(get(handles.g_editText,'String'));
alpha = str2num(get(handles.alpha_editText,'String'));
X_c = str2num(get(handles.X_c_editText,'String'));
b = str2num(get(handles.b_editText,'String'));
num_of_cycles = str2num(get(handles.num_of_cycles_editText,'String'));

switch get(handles.model_popupmenu,'Value')
    case 1
        model=0; % Linear Ion Drift
    case 2
        model=1; % Simmons Tunnel Barrier
    case 3
        model=2; % Team
    case 4
        model=3; % Nonlinear Ion Drift 
    otherwise
end

switch get(handles.window_popupmenu,'Value')
    case 1
        win=0; %ideal window
    case 2
        win=1; %Jogelkar window
    case 3
        win=2; %Biolek window
    case 4
        win=3; %Prodromakis window
    case 5
        win=4; %Kvatinsky window only recognized for Team model
    otherwise
end

%% Linear Ion Drift model W plot
if (model==0)  
 
tspan=[0 num_of_cycles/freq];                       %%time length of the simulation
points=2e5;                              %%number of sampling points
W0=w_init*D;                            %define the initial value of W
tspan_vector = linspace(tspan(1),tspan(2),points);         % Create vector of initial values
I = amp*sin(freq*2*pi*tspan_vector);                                       %%can also use square wave generated by : (square(tspan_vector));
W=zeros(size((tspan_vector)));
W_dot=zeros(size((tspan_vector)));
delta_t=tspan_vector(2)-tspan_vector(1);                        %%define the step size

W(1)=W0;                                                                                                      %% initiliaze the first W vetor elemnt to W0 - the initial condition
for i=2:length(tspan_vector)
    % case this is an ideal window
    if (((win==0) || (win==4)) && ((abs (I(i))) >= (V_t/(Ron*W(i-1)/D+Roff*(1-W(i-1)/D))))) 
        W_dot(i)=I(i)*(Ron*uV/D);
        W(i)=W(i-1)+W_dot(i)*delta_t;
    elseif ((win==0) && ((abs(I(i)))  < (V_t/(Ron*W(i-1)/D+Roff*(1-W(i-1)/D)))))
        W(i)=W(i-1);
        W_dot(i)=0;
        
    end
    
    % case this is Jogelkar window
    if ((win==1) && ((abs(I(i))) >= (V_t/(Ron*W(i-1)/D+Roff*(1-W(i-1)/D)))))
        W_dot(i)=I(i)*(Ron*uV/D);
        W(i)=W(i-1)+W_dot(i)*delta_t*(1-(2*W(i-1)/D-1)^(2*P_coeff));%+1e-18*sign(I(i));
    elseif ((win==1) && ((abs (I(i)) ) < (V_t/(Ron*W(i-1)/D+Roff*(1-W(i-1)/D)))))
        W(i)=W(i-1);
        W_dot(i)=0;

    end
    
        % case this is Biolek window
    if ((win==2) && ((abs(I(i))) >= (V_t/(Ron*W(i-1)/D+Roff*(1-W(i-1)/D)))))
        W_dot(i)=I(i)*(Ron*uV/D);
        W(i)=W(i-1)+W_dot(i)*delta_t*(1-(W(i-1)/D-heaviside(-I(i)))^(2*P_coeff));
    elseif ((win==2) && ((abs(I(i))) < (V_t/(Ron*W(i-1)/D+Roff*(1-W(i-1)/D)))))
        W(i)=W(i-1);
        W_dot(i)=0;
    end
 
        % case this is Prodromakis window
    if ((win==3) && ((abs(I(i))) >= (V_t/(Ron*W(i-1)/D+Roff*(1-W(i-1)/D)))))
        W_dot(i)=I(i)*(Ron*uV/D);
        W(i)=W(i-1)+W_dot(i)*delta_t*(J*(1-((W(i-1)/D-0.5)^2+0.75)^P_coeff));
    elseif ((win==3) && ((abs(I(i))) < (V_t/(Ron*W(i-1)/D+Roff*(1-W(i-1)/D)))))
        W(i)=W(i-1);
        W_dot(i)=0;
    end
    
  % correct the w vector according to bounds [0 D]
    if W(i) < 0
        W(i) = 0;
        W_dot(i)=0;
    elseif W(i) > D
        W(i) = D;
        W_dot(i)=0;
    end
end

R=Ron*W/D+Roff*(1-W/D); %%this parameter might be useful for debug
V= R.*I; 
plot(tspan_vector,W/D,'r');
title('W/D as func of time');
xlabel('time[sec]');
legend('W/D')
end

%%  Simmons Tunnel Barrier model X plot
if (model==1)

points=2e5;
tspan=[0 num_of_cycles/freq];
t = linspace(tspan(1),tspan(2),points);
curr = amp*sin(freq*2*pi*t);
X=zeros(1,points);
X_dot=zeros(1,points);
delta_t=t(2)-t(1);

X(1)=w_init*D;
      
for i=2:(length(t))
          if curr(i)> 0
               X_dot(i)=c_off*sinh(curr(i)/i_off)*exp(-exp((X(i-1)-a_off)/X_c-abs(curr(i))/b)-X(i-1)/X_c);
          else
              X_dot(i)=c_on*sinh(curr(i)/i_on)*exp(-exp(-(X(i-1)-a_on)/X_c-abs(curr(i))/b)-X(i-1)/X_c);
          end
          
         X(i)=X(i-1)+delta_t*X_dot(i);
         
    if (X(i) < 0)
        X(i) = 0;
        X_dot(i)=0;
    elseif (X(i) > D)
        X(i) = D;
        X_dot(i)=0;
    end
end

R=Roff.*X./D+Ron.*(1-X./D);
V=R.*curr;

plot(t,X/D,'r');
title('X/D  as func of time');
xlabel('time[sec]');
legend('X/D');
end

%%  Team model X plot

if (model==2)

points=2e5;
tspan=[0 num_of_cycles/freq];
t = linspace(tspan(1),tspan(2),points);
curr = amp*sin(freq*2*pi*t);
X=zeros(1,points);
X_dot=zeros(1,points);
delta_t=t(2)-t(1);

X(1)=w_init*D;
     
for i=2:(length(t))
    % case this is an ideal window
    if (win == 0) 
          if (curr(i) > 0) && (curr(i) > i_off)
               X_dot(i)=k_off*(curr(i)/i_off-1)^alpha_off;
               X(i)=X(i-1)+delta_t*X_dot(i);
          elseif (curr(i) <= 0) && (curr(i) < i_on)
               X_dot(i)=k_on*(curr(i)/i_on-1)^alpha_on;
               X(i)=X(i-1)+delta_t*X_dot(i);
          else
               X(i)=X(i-1);
               X_dot(i)=0;
          end
    end
    
    % case this is a Jogelkar window
    if (win == 1) 
          if (curr(i) > 0) && (curr(i) > i_off)
               X_dot(i)=k_off*(curr(i)/i_off-1)^alpha_off;
               X(i)=X(i-1)+delta_t*X_dot(i).*(1-(2*X(i-1)/D-1)^(2*P_coeff));
          elseif (curr(i) <= 0) && (curr(i) < i_on)
               X_dot(i)=k_on*(curr(i)/i_on-1)^alpha_on;
               X(i)=X(i-1)+delta_t*X_dot(i).*(1-(2*X(i-1)/D-1)^(2*P_coeff));
          else
               X(i)=X(i-1);
               X_dot(i)=0;
          end
    end
    
    % case this is Biolek window
    if (win == 2) 
          if (curr(i) > 0) && (curr(i) > i_off)
               X_dot(i)=k_off*(curr(i)/i_off-1)^alpha_off;
               X(i)=X(i-1)+delta_t*X_dot(i).*(1-(1-X(i-1)/D-heaviside(curr(i)))^(2*P_coeff));
          elseif (curr(i) <= 0) && (curr(i) < i_on)
               X_dot(i)=k_on*(curr(i)/i_on-1)^alpha_on;
               X(i)=X(i-1)+delta_t*X_dot(i).*(1-(1-X(i-1)/D-heaviside(curr(i)))^(2*P_coeff));
          else
               X(i)=X(i-1);
               X_dot(i)=0;
          end
    end
    
    % case this is Prodromakis window
    if (win == 3)
          if (curr(i) > 0) && (curr(i) > i_off)
               X_dot(i)=k_off*(curr(i)/i_off-1)^alpha_off;
               X(i)=X(i-1)+delta_t*X_dot(i).*(J*(1-((X(i-1)/D-0.5)^2+0.75)^P_coeff));
          elseif (curr(i) <= 0) && (curr(i) < i_on)
               X_dot(i)=k_on*(curr(i)/i_on-1)^alpha_on;
               X(i)=X(i-1)+delta_t*X_dot(i).*(J*(1-((X(i-1)/D-0.5)^2+0.75)^P_coeff));
          else
               X(i)=X(i-1);
               X_dot(i)=0;
          end
    end
    
    % case this is Kvatinsky window
    if (win == 4) 
          if (curr(i) > 0) && (curr(i) > i_off)
               X_dot(i)=k_off*(curr(i)/i_off-1)^alpha_off*exp(-exp((X(i-1)-a_off)/X_c));
               X(i)=X(i-1)+delta_t*X_dot(i);
          elseif (curr(i) <= 0) && (curr(i) < i_on)
               X_dot(i)=k_on*(curr(i)/i_on-1)^alpha_on*exp(-exp(-(X(i-1)-a_on)/X_c));
               X(i)=X(i-1)+delta_t*X_dot(i);
          else
               X(i)=X(i-1);
               X_dot(i)=0;
          end
    end
     
    if (X(i) < 0)
        X(i) = 0;
        X_dot(i)=0;
    elseif (X(i) > D)
        X(i) = D;
        X_dot(i)=0;
    end


end


plot(t,X/D);
title('X/D as func of time');
xlabel('time[sec]');
legend('X/D')
    
end
%%  Nonlinear Ion Drift model X plot
if (model==3)
points=40000;
tspan=[0 num_of_cycles/freq];
t = linspace(tspan(1),tspan(2),points);
delta_t = t(2) - t(1);
V = amp*sin(freq*2*pi*t);
W=zeros(size((t)));
W_dot=zeros(size((t)));
I=zeros(size((t)));
W(1) = w_init;
curr(1)=0;

for i=2:points
    % case this is an ideal window
    if ((win==0) || (win==4))
        W_dot(i)=a*V(i)^q;
        W(i)=W(i-1)+W_dot(i)*delta_t;  
    end
    
    % case this is Jogelkar window
    if (win==1)
        W_dot(i)=a*V(i)^q;
        W(i)=W(i-1)+W_dot(i)*delta_t*(1-(2*W(i-1)-1)^(2*P_coeff));
    end
    
     % case this is Biolek window
    if (win==2)
        W_dot(i)=a*V(i)^q;
        W(i)=W(i-1)+W_dot(i)*delta_t*(1-(W(i-1)-heaviside(-V(i-1)))^(2*P_coeff));
   end
 
    % case this is Prodromakis window
    if (win==3)
        W_dot(i)=a*V(i)^q;
        W(i)=W(i-1)+W_dot(i)*delta_t*(J*(1-((W(i-1)-0.5)^2+0.75)^P_coeff));
    end
    
    % correct the w vector according to bounds [0 D]
    if W(i) < 0
        W(i) = 0;
        W_dot(i)=0;
    elseif W(i) > 1
        W(i) = 1;
        W_dot(i)=0;
    end
    
  curr(i)=W(i)^n*beta*sinh(alpha*V(i))+c*(exp(g*V(i))-1);
    
    
end

plot(t,W,'r');
title('W/D');
xlabel('time[sec]');
legend('W/D')

end


guidata(hObject, handles); %updates the handles


% --- Executes on button press in plotIV_pushbutton.
function plotIV_pushbutton_Callback(hObject, eventdata, handles)
% hObject    handle to plotIV_pushbutton (see GCBO)
% eventdata  reserved - to be defined in a future version of MATLAB
% handles    structure with handles and user data (see GUIDATA)

%selects axes2 as the current axes, so that
%Matlab knows where to plot the data
axes(handles.axes2)

amp = str2num(get(handles.amplitude_editText,'String'));
Roff = str2num(get(handles.Roff_editText,'String'));
Ron = str2num(get(handles.Ron_editText,'String'));
freq = str2num(get(handles.frequency_editText,'String'));
D = str2num(get(handles.D_editText,'String'));
uV =  str2num(get(handles.uV_editText,'String'));
V_t = str2num(get(handles.V_t_editText,'String'));
P_coeff = str2num(get(handles.p_coeff_editText,'String'));
w_init = str2num(get(handles.W0_editText,'String'));
J = str2num(get(handles.J_editText,'String'));
a_on = str2num(get(handles.a_on_editText,'String'));
a_off = str2num(get(handles.a_off_editText,'String'));
c_on = str2num(get(handles.c_on_editText,'String'));
c_off = str2num(get(handles.c_off_editText,'String'));
alpha_on = str2num(get(handles.alpha_on_editText,'String'));
alpha_off = str2num(get(handles.alpha_off_editText,'String'));
k_on = str2num(get(handles.k_on_editText,'String'));
k_off = str2num(get(handles.k_off_editText,'String'));
i_on = str2num(get(handles.i_on_editText,'String'));
i_off = str2num(get(handles.i_off_editText,'String'));
x_on = str2num(get(handles.x_on_editText,'String'));
x_off = str2num(get(handles.x_off_editText,'String'));
beta = str2num(get(handles.beta_editText,'String'));
a = str2num(get(handles.a_editText,'String'));
c = str2num(get(handles.c_editText,'String'));
n = str2num(get(handles.n_editText,'String'));
q = str2num(get(handles.q_editText,'String'));
g = str2num(get(handles.g_editText,'String'));
alpha = str2num(get(handles.alpha_editText,'String'));
X_c = str2num(get(handles.X_c_editText,'String'));
b = str2num(get(handles.b_editText,'String'));
num_of_cycles = str2num(get(handles.num_of_cycles_editText,'String'));

switch get(handles.model_popupmenu,'Value')
    case 1
        model=0; % Linear Ion Drift
    case 2
        model=1; % Simmons Tunnel Barrier
    case 3
        model=2; % Team
    case 4
        model=3; % Nonlinear Ion Drift
    otherwise
end

switch get(handles.window_popupmenu,'Value')
    case 1
        win=0; %ideal window
    case 2
        win=1; %Jogelkar window
    case 3
        win=2; %Biolek window
    case 4
        win=3; %Prodromakis window
    case 5
        win=4; %Kvatinsky window only recognized for Team model
    otherwise
end

switch get(handles.iv_popupmenu,'Value')
    case 1
        iv=0; %V=IR
    case 2
        iv=1; %V=I*exp(...)
    otherwise
end

%% Linear Ion Drift model I-V plot
if (model==0)  
 
tspan=[0 num_of_cycles/freq];                       %%time length of the simulation
points=2e5;                              %%number of sampling points
W0=w_init*D;                            %define the initial value of W
tspan_vector = linspace(tspan(1),tspan(2),points);         % Create vector of initial values
I = amp*sin(freq*2*pi*tspan_vector);                   %%can also use square wave generated by : (square(tspan_vector));
W=zeros(size((tspan_vector)));
W_dot=zeros(size((tspan_vector)));
delta_t=tspan_vector(2)-tspan_vector(1);                        %%define the step size

W(1)=W0;                                                 %% initiliaze the first W vetor elemnt to W0 - the initial condition
for i=2:length(tspan_vector)
    % case this is an ideal window
    if (((win==0) || (win==4)) && ((abs (I(i))) >= (V_t/ (Ron*W(i-1)/D+Roff*(1-W(i-1)/D))))) 
        W_dot(i)=I(i)*(Ron*uV/D);
        W(i)=W(i-1)+W_dot(i)*delta_t;
    elseif ((win==0) && ((abs(I(i)))  < (V_t/ (Ron*W(i-1)/D+Roff*(1-W(i-1)/D)))))
        W(i)=W(i-1);
        W_dot(i)=0;
        
    end
    
    % case this is Jogelkar window
    if ((win==1) && ((abs(I(i))) >= (V_t/(Ron*W(i-1)/D+Roff*(1-W(i-1)/D)))))
        W_dot(i)=I(i)*(Ron*uV/D);
        W(i)=W(i-1)+W_dot(i)*delta_t*(1-(2*W(i-1)/D-1)^(2*P_coeff));%%+1e-18*sign(I(i));
    elseif ((win==1) && ((abs (I(i)) ) < (V_t/ (Ron*W(i-1)/D+Roff*(1-W(i-1)/D)))))
        W(i)=W(i-1);
        W_dot(i)=0;

    end
    
        % case this is Biolek window
    if ((win==2) && ((abs(I(i))) >= (V_t/(Ron*W(i-1)/D+Roff*(1-W(i-1)/D)))))
        W_dot(i)=I(i)*(Ron*uV/D);
        W(i)=W(i-1)+W_dot(i)*delta_t*(1-(W(i-1)/D-heaviside(-I(i)))^(2*P_coeff));
    elseif ((win==2) && ((abs(I(i))) < (V_t/(Ron*W(i-1)/D+Roff*(1-W(i-1)/D)))))
        W(i)=W(i-1);
        W_dot(i)=0;
    end
 
        % case this is Prodromakis window
    if ((win==3) && ((abs(I(i))) >= (V_t/(Ron*W(i-1)/D+Roff*(1-W(i-1)/D)))))
        W_dot(i)=I(i)*(Ron*uV/D);
        W(i)=W(i-1)+W_dot(i)*delta_t*(J*(1-((W(i-1)/D-0.5)^2+0.75)^P_coeff));
    elseif ((win==3) && ((abs(I(i))) < (V_t/(Ron*W(i-1)/D+Roff*(1-W(i-1)/D)))))
        W(i)=W(i-1);
        W_dot(i)=0;
    end
    
  % correct the w vector according to bounds [0 D]
    if W(i) < 0
        W(i) = 0;
        W_dot(i)=0;
    elseif W(i) > D
        W(i) = D;
        W_dot(i)=0;
    end
end

R=Ron*W/D+Roff*(1-W/D); %%this parameter might be useful for debug
V= R.*I; 
plot(V(20e3:end),I(20e3:end));
title('I-V curve');
xlabel('V[volt]');
ylabel('I[amp]');

end

%%  Simmons Tunnel Barrier model I-V plot
if (model==1)

points=2e5;
tspan=[0 num_of_cycles/freq];
t = linspace(tspan(1),tspan(2),points);
curr = amp*sin(freq*2*pi*t);
X=zeros(1,points);
X_dot=zeros(1,points);
delta_t=t(2)-t(1);

X(1)=w_init*D;
      
for i=2:(length(t))
          if curr(i)> 0
               X_dot(i)=c_off*sinh(curr(i)/i_off)*exp(-exp((X(i-1)-a_off)/X_c-abs(curr(i))/b)-X(i-1)/X_c);
          else
              X_dot(i)=c_on*sinh(curr(i)/i_on)*exp(-exp(-(X(i-1)-a_on)/X_c-abs(curr(i))/b)-X(i-1)/X_c);
          end
          
         X(i)=X(i-1)+delta_t*X_dot(i);
         
    if (X(i) < 0)
        X(i) = 0;
        X_dot(i)=0;
    elseif (X(i) > D)
        X(i) = D;
        X_dot(i)=0;
    end
end

R=Roff.*X./D+Ron.*(1-X./D);
V=R.*curr;

plot(V(20e3:end),curr(20e3:end));
title('I-V curve');
xlabel('V[volt]');
ylabel('I[amp]');

end

%%  Team model I-V plot

if (model==2)
points=2e5;
tspan=[0 num_of_cycles/freq];
t = linspace(tspan(1),tspan(2),points);
curr = amp*sin(freq*2*pi*t);
X=zeros(1,points);
X_dot=zeros(1,points);
delta_t=t(2)-t(1);

X(1)=w_init*D;
     
for i=2:(length(t))
    % case this is an ideal window
    if (win == 0) 
          
          if (curr(i) > 0) && (curr(i) > i_off)
               X_dot(i)=k_off*(curr(i)/i_off-1)^alpha_off;
               X(i)=X(i-1)+delta_t*X_dot(i);
          elseif (curr(i) <= 0) && (curr(i) < i_on)
               X_dot(i)=k_on*(curr(i)/i_on-1)^alpha_on;
               X(i)=X(i-1)+delta_t*X_dot(i);
          else
               X(i)=X(i-1);
               X_dot(i)=0;
          end
    end
    
    % case this is Jogelkar window
    if (win == 1) 
          if (curr(i) > 0) && (curr(i) > i_off)
               X_dot(i)=k_off*(curr(i)/i_off-1)^alpha_off;
               X(i)=X(i-1)+delta_t*X_dot(i).*(1-(2*X(i-1)/D-1)^(2*P_coeff));
          elseif (curr(i) <= 0) && (curr(i) < i_on)
               X_dot(i)=k_on*(curr(i)/i_on-1)^alpha_on;
               X(i)=X(i-1)+delta_t*X_dot(i).*(1-(2*X(i-1)/D-1)^(2*P_coeff));
          else
               X(i)=X(i-1);
               X_dot(i)=0;
          end
    end
    
    % case this is Biolek window
    if (win == 2) 
          if (curr(i) > 0) && (curr(i) > i_off)
               X_dot(i)=k_off*(curr(i)/i_off-1)^alpha_off;
               X(i)=X(i-1)+delta_t*X_dot(i).*(1-(1-X(i-1)/D-heaviside(curr(i)))^(2*P_coeff));
          elseif (curr(i) <= 0) && (curr(i) < i_on)
               X_dot(i)=k_on*(curr(i)/i_on-1)^alpha_on;
               X(i)=X(i-1)+delta_t*X_dot(i).*(1-(1-X(i-1)/D-heaviside(curr(i)))^(2*P_coeff));
          else
               X(i)=X(i-1);
               X_dot(i)=0;
          end
    end
    
    % case this is Prodromakis window
    if (win == 3)
          if (curr(i) > 0) && (curr(i) > i_off)
               X_dot(i)=k_off*(curr(i)/i_off-1)^alpha_off;
               X(i)=X(i-1)+delta_t*X_dot(i).*(J*(1-((X(i-1)/D-0.5)^2+0.75)^P_coeff));
          elseif (curr(i) <= 0) && (curr(i) < i_on)
               X_dot(i)=k_on*(curr(i)/i_on-1)^alpha_on;
               X(i)=X(i-1)+delta_t*X_dot(i).*(J*(1-((X(i-1)/D-0.5)^2+0.75)^P_coeff));
          else
               X(i)=X(i-1);
               X_dot(i)=0;
          end
    end
    
    % case this is Kvatinsky window
    if (win == 4)
          if (curr(i) > 0) && (curr(i) > i_off)
               X_dot(i)=k_off*(curr(i)/i_off-1)^alpha_off*exp(-exp((X(i-1)-a_off)/X_c));
               X(i)=X(i-1)+delta_t*X_dot(i);
          elseif (curr(i) <= 0) && (curr(i) < i_on)
               X_dot(i)=k_on*(curr(i)/i_on-1)^alpha_on*exp(-exp(-(X(i-1)-a_on)/X_c));
               X(i)=X(i-1)+delta_t*X_dot(i);
          else
               X(i)=X(i-1);
               X_dot(i)=0;
          end
    end
     
    if (X(i) < 0)
        X(i) = 0;
        X_dot(i)=0;
    elseif (X(i) > D)
        X(i) = D;
        X_dot(i)=0;
    end


end

    if (iv == 0) %case I-V relation is linear
       R=Roff.*X./D+Ron.*(1-X./D);
       V=R.*curr;
    else %case the I-V relation is nonlinear
       lambda = log(Roff/Ron);
       V=Ron*curr.*exp(lambda*(X-x_on)/(x_off-x_on));
    end
 
plot(V(20e3:end),curr(20e3:end));
title('I-V curve');
xlabel('V[volt]');
ylabel('I[amp]');

end

%%  Nonlinear Ion Drift model I-V plot
if (model==3)
points=40000;
tspan=[0 num_of_cycles/freq];
t = linspace(tspan(1),tspan(2),points);
delta_t = t(2) - t(1);
V = amp*sin(freq*2*pi*t);
W=zeros(size((t)));
W_dot=zeros(size((t)));
curr=zeros(size((t)));
W(1) = w_init;

for i=2:points
    % case this is an ideal window
    if ((win==0) || (win==4))
        W_dot(i)=a*V(i)^q;
        W(i)=W(i-1)+W_dot(i)*delta_t;  
    end
    
    % case this is Jogelkar window
    if (win==1)
        W_dot(i)=a*V(i)^q;
        W(i)=W(i-1)+W_dot(i)*delta_t*(1-(2*W(i-1)-1)^(2*P_coeff));
    end
    
     % case this is Biolek window
    if (win==2)
        W_dot(i)=a*V(i)^q;
        W(i)=W(i-1)+W_dot(i)*delta_t*(1-(W(i-1)-heaviside(-V(i-1)))^(2*P_coeff));
   end
 
    % case this is Prodromakis window
    if (win==3)
        W_dot(i)=a*V(i)^q;
        W(i)=W(i-1)+W_dot(i)*delta_t*(J*(1-((W(i-1)-0.5)^2+0.75)^P_coeff));
    end
    
    % correct the w vector according to bounds [0 D]
    if W(i) < 0
        W(i) = 0;
        W_dot(i)=0;
    elseif W(i) > 1
        W(i) = 1;
        W_dot(i)=0;
    end
    
  curr(i)=W(i)^n*beta*sinh(alpha*V(i))+c*(exp(g*V(i))-1);
    
    
end

plot(V(20e3:end),curr(20e3:end));
title('I-V curve');
xlabel('V[volt]');
ylabel('I[amp]');

end
guidata(hObject, handles); %updates the handles


% --- Executes on button press in clearplots_pushbutton.
function clearplots_pushbutton_Callback(hObject, eventdata, handles)
% hObject    handle to clearplots_pushbutton (see GCBO)
% eventdata  reserved - to be defined in a future version of MATLAB
% handles    structure with handles and user data (see GUIDATA)
%these two lines of code clears both axes
cla(handles.axes1,'reset')
cla(handles.axes2,'reset')
guidata(hObject, handles); %updates the handles

