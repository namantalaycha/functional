#!/bin/sh; C:/Program\ Files/Git/usr/bin/sh.exe

SMTPFROM=naman.talaycha@myntra.com
SMTPTO=namanfrost10@gmail.com
SMTPSERVER=smtp.googlemail.com:587
SMTPUSER=naman.talaycha@myntra.com
SMTPPASS=zpbwaqijzukxwkee
MESSAGEBODY="uegeu eieigiger"
SUBJECT="testing"


sendEmail -f $SMTPFROM -t $SMTPTO -u $SUBJECT -m $MESSAGEBODY -s $SMTPSERVER -xu $SMTPUSER -xp $SMTPPASS

num='ok'

if [ "$num" ==  'ok' ] 
then 
        echo "Test do not pass"
        exit 1
fi
