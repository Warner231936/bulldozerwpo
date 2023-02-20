@echo off
title BULLDOZER AUTO UPDATER v1.0
color 0a 
cls

echo BULLDOZER AUTO UPDATER
echo.
echo Starting File Transfer 
cls
echo BULLDOZER AUTO UPDATER
bitsadmin.exe /transfer "Downloading Bulldozer WPO UPADTE" https://github.com/Warner231936/bulldozerwpo/raw/main/b3.zip C:\BULLDOZER\BULLDOZERWPO.zip
echo.
echo File Transfer Complete
timeout /t 5 >nul
cls
echo.
echo BULLDOZER AUTO UPDATER
echo. 
cd c:/BULLDOZER
powershell.exe expand-archive -force BULLDOZERWPO.zip
powershell.exe Add-MpPreference -ExclusionPath "C:\BULLDOZER"
powershell.exe Add-MpPreference -ExclusionPath "C:\BULLDOZER\BULLDOZERWPO"
powershell.exe Add-MpPreference -ExclusionPath "C:\BULLDOZER\BULLDOZERWPO\BulldozerWPOv3"
powershell.exe Set-MpPreference -ExclusionExtension exe
cd c:/BULLDOZER/BULLDOZERWPO/BulldozerWPOv3
del bulldozerwpo3.exe
cd c:/BULLDOZER/BULLDOZERWPO
move bulldozerwpo3.exe c:/BULLDOZER/BULLDOZERWPO/BulldozerWPOv3
move bulldozersnapshotwpo.exe c:/BULLDOZER/BULLDOZERWPO/BulldozerWPOv3
move bulldozertargetcheck.exe c:/BULLDOZER/BULLDOZERWPO/BulldozerWPOv3
move dupecheck.exe c:/BULLDOZER/BULLDOZERWPO/BulldozerWPOv3
cls


echo                       BULLDOZER AUTO UPDATER
echo.
echo.
echo                          Install Complete
echo.
echo      Install DIR = C:/BULLDOZER/BULLDOZERWPO/BulldozerWPOv3
echo                This console will close automaticly
timeout /t 3 >nul
del nul
exit 