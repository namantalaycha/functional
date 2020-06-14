
import smtplib

server = smtplib.SMTP_SSL("smtp.gmail.com" , 465)
server.login("naman.talaycha@myntra.com" , "zpbwaqijzukxwkee")
server.sendmail("naman.talaycha@myntra.com" , "namanfrost10@gmail.com" , "Testing Successfull")

server.quit()