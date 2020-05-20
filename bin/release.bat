@echo off
echo ---
echo default_process_types:
IF "%HBP_APP_ARCHITECTURE%" == "x86" (
echo   web: .cloudfoundry\hwc_x86.exe
) ELSE (
echo   web: .cloudfoundry\hwc.exe
)
