#! /bin/sh

hostname >> /testpv/podinfo.txt
echo "-------------------cpuinfo---------------" >> /testpv/podinfo.txt
cat /proc/cpuinfo >> /testpv/podinfo.txt
echo "-------------------meminfo---------------" >> /testpv/podinfo.txt
cat /proc/meminfo >> /testpv/podinfo.txt