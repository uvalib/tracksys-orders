# run application

cd bin; ./tsorders -dbhost $DBHOST -dbport $DBPORT -dbname $DBNAME -dbuser $DBUSER -dbpass $DBPASS -smtphost $SMPT_HOST -smtpport $SMPT_PORT -smtpuser $SMPT_USER -smtppass $SMPT_PASS -smtpsender $SMPT_SENDER

# return the status
exit $?

#
# end of file
#
