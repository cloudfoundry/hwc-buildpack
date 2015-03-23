SET BIN=%CD%\bin
SET PATH=%BIN%;%PATH%

:: enable some features
SET dism=%WINDIR%\SysNative\dism.exe
%dism% /online /Enable-Feature /FeatureName:IIS-WebServer /All /NoRestart
%dism% /online /Enable-Feature /FeatureName:IIS-WebSockets /All /NoRestart
%dism% /online /Enable-Feature /FeatureName:Application-Server-WebServer-Support /FeatureName:AS-NET-Framework /All /NoRestart
%dism% /online /Enable-Feature /FeatureName:IIS-HostableWebCore /All /NoRestart
 
del /F /Q windows_app_lifecycle.tgz
rmdir /S /Q packages
nuget restore || exit /b 1

MSBuild WindowsCircus.sln /t:Rebuild /p:Configuration=Release || exit /b 1
packages\nspec.0.9.68\tools\NSpecRunner.exe Builder.Tests\bin\Release\BuilderTests.dll || exit /b 1
packages\nspec.0.9.68\tools\NSpecRunner.exe Launcher.Tests\bin\Release\LauncherTests.dll || exit /b 1
packages\nspec.0.9.68\tools\NSpecRunner.exe WebAppServer.Tests\bin\Release\WebAppServer.Tests.dll || exit /b 1
bsdtar -czvf windows_app_lifecycle.tgz --exclude log -C Builder\bin . -C ..\..\Launcher\bin . -C ..\..\Healthcheck\bin . -C ..\..\WebAppServer\bin . || exit /b 1
