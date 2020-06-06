

FILE_PATH=/home/deepesh/ShekharScripting/MyTest.json
KafkaBroker="localhost:9092,localhost:9093"
URL="abc.com"
echo $KafkaBroker
echo $URL

KafkaBrokers=$(echo $KafkaBroker | sed 's/,/","/g')

echo $KafkaBrokers

var_replace() {
  key=$1
  value=$2
  awk -v val="$value" "/$key/{\$2=val}1" $FILE_PATH > tmp_file && mv tmp_file $FILE_PATH
  #awk -v FS=":" '{print $1"\t"$2}' $FILE_PATH
}

#var_replace Brokers "xyzshekhar1"
var_replace Brokers "[\"$KafkaBrokers\"],"
var_replace URL "\"$URL\","
echo "Number of paramters are"
echo $#
