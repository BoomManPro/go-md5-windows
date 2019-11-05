@echo off

reg add "HKEY_CLASSES_ROOT\*\Shell\md5File"  /d "" /f

reg add "HKEY_CLASSES_ROOT\*\Shell\md5File\command"  /d "\"%~dp0md5.exe\"  \"%%1\"" /f
