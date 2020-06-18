#!/bin/bash

touch .git/hooks/pre-commit
chmod +x .git/hooks/pre-commit
echo "#!/bin/bash" > .git/hooks/pre-commit
echo "./runTest.sh" >> .git/hooks/pre-commit
echo "./checkFail.sh" >> .git/hooks/pre-commit

echo "python3 slackbot.py" >> .git/hooks/pre-commit

sed -e 's/:[^:\/\/]/=/g;s/$//g;s/ *=/=/g' configure.yaml >> .git/hooks/pre-commit

echo "SMTPSERVER=smtp.googlemail.com:587" >> .git/hooks/pre-commit

echo "sendEmail -f \$SMTPFROM -t \$SMTPTO -u \$SUBJECT -o message-file=temp/text.txt -s \$SMTPSERVER -xu \$SMTPUSER -xp \$SMTPPASS" >> .git/hooks/pre-commit

echo "rm -r temp/text.txt " >> .git/hooks/pre-commit