#!/bin/zsh
set -x

for f in $( find . -name "*.proto"  -wholename  './k8s.io/*'  | xargs grep go_package | grep -v "= \"k8s" | awk -F':' '{print $1}');do
#  package=$(cat $f | grep go_package | grep -v "= \"k8s" | awk -F':' '{print $1}' | awk -F"\/" '{pack=$2;for(i=3; i<=NF-1; i++){pack=pack"\/" $i} print pack;}')
  package=$(echo $f |  awk -F':' '{print $1}' | awk -F"\/" '{pack=$2;for(i=3; i<=NF-1; i++){pack=pack"\\/" $i} print pack;}')
  suffix=$(echo $f | awk -F"\/" '{print $(NF-1)}')
  sed -i '' 's/\"'${suffix}'\"/\"'${package}'\"/g' $f
done