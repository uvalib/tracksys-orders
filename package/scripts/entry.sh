# run application
# set blank options variables
SMTP_USER_OPT=""
SMTP_PASS_OPT=""

# SMTP username
if [ -n "${SMPT_USER}" ]; then
   SMTP_USER_OPT="-smtpuser ${SMPT_USER}"
fi

# SMTP password
if [ -n "${SMPT_PASS}" ]; then
   SMTP_PASS_OPT="-smtppass ${SMPT_PASS}"
fi

cd bin; ./tsorders            \
   -dbhost $DBHOST            \
   -dbport $DBPORT            \
   -dbname $DBNAME            \
   -dbuser $DBUSER            \
   -dbpass $DBPASS            \
   -smtphost $SMPT_HOST       \
   -smtpport $SMPT_PORT       \
   -smtpsender $SMPT_SENDER   \
   ${SMTP_USER_OPT}           \
   ${SMTP_PASS_OPT}


# return the status
exit $?

#
# end of file
#
