ARGS=$@

function on_done() {
  echo "Job dispached."
  exit 0
}

if [ "$1" = '-h' ] || [ "$1" = '--help' ]
then
  python3 remote/remote.py -h
elif [[ "$1" = *"-k"* ]];
then
  NARGS=$#
  if [[ $NARGS = 1 && ($1 = "-k" || $1 = "--kill") ]];
  then
    echo "Killing all remote jobs..."
    python3 remote/remote.py -k > /dev/null 2>&1 &
  else
    echo "Killing single remote job..."
    python3 remote/remote.py $ARGS
  fi
else
  pid=$$
  echo "Dispaching job: $ARGS"
  echo "PID: $pid"

  # Wait for SIGUR1 signal
  trap "on_done" SIGUSR1

  nohup python3 remote/remote.py -notify=$pid $ARGS > /dev/null 2>&1 &
  # python3 remote/remote.py -notify=$pid $ARGS

  while :; do
    sleep 1
  done
fi
