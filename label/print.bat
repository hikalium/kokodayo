SET SPC_PATH="C:\Program Files (x86)\KING JIM\TEPRA SPC10\SPC10.exe"
SET CWD=%~dp0
SET TPE_PATH=%CWD%boxes.tpe
rem CSV_PATH should end with .csv
SET CSV_PATH=%CWD%entries.csv
rem SET FIRST_IDX=%1
rem SET LAST_IDX=%2
set /p FIRST_IDX=First index [number]?: 
set /p LAST_IDX=Last index[number]?: 
copy /Y nul: %CSV_PATH%
for /l %%i in (%FIRST_IDX%,1,%LAST_IDX%) do (
    rem repeat 2 times to print together
    echo %%i, b/%%i, http://kokodayo/b/%%i >> %CSV_PATH%
    echo %%i, b/%%i, http://kokodayo/b/%%i >> %CSV_PATH%
)
type %CSV_PATH%
echo %TPE_PATH%
%SPC_PATH% /p "%TPE_PATH%,%CSV_PATH%,1"
