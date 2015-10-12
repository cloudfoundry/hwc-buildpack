@echo off
setlocal enabledelayedexpansion

del Bean.txt

set argCount=0
for %%x in (%*) do (
   set /A argCount+=1
   set "argVec[!argCount!]=%%~x"
)

for /L %%i in (1,1,%argCount%) do (
  echo "!argVec[%%i]!" >> Bean.txt
)

echo This is STDOUT
echo This is STDERR >&2
