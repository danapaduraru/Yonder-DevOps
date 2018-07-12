```bash
#Scriptul de bash care logheaza consumul de resurse pe sistem.  
  
#!/bin/bash  
  
if [[ -f "/var/log/systemusage/process.log" ]]  
then   
	echo "Fisierul process.log exista"  
else  
	touch /var/log/systemusage/process.log;  
	echo "Fisierul process.log a fost creat";  
fi  
#afisare data sub formatul din exemplu:  
printf "$(date +"%d %b %Y %T:") \n" >>/var/log/systemusage/process.log  
#top-afiseaza procesele, b + n1 il face stabil dupa 1 refresh  
#head-afiseaza primele 12 linii  
#tail-afiseaza ultimele 5 linii  
#awk face operatii pe linii - aici afiseaza doar liniile 12, 1, 9  
top -b -n1 -o%CPU | head -n12 | tail -n5 | awk '{print  $12"\t"$1"\t"$9}' >>/var/log/systemusage/process.log  
  
  
if [[ -f "var/log/systemusage/disk.log" ]]  
then   
	echo "Fisierul disk.log exista"  
else  
	touch /var/log/systemusage/disk.log;  
	echo "Fisierul disk.log a fost creat";  
fi  
printf "$(date +"%d %b %Y %T:") \n" >>/var/log/systemusage/disk.log  
#df - afiseaza datele referitoare la memorie  
#tr - concateneaza spatiile astfel incat sa avem doar unul, pentru a face sort  
#sort - sorteaza descrescator dupa field 5, separator " ",  
#grep - va afisa tot in afara de ce este tempfs  
#tail - afiseaza ultimele linii in afara de prima (incepe de la a doua linie +2)  
df | tr -s ' ' | sort -r -t" " -k5 | grep tmpfs -v | tail -n +2 | awk '{print  $1"\t"$5"\t"$6}' >>/var/log/systemusage/disk.log  
  
if [[ -f "var/log/systemusage/memory.log" ]]  
then   
	echo "Fisierul memory.log exista"  
else  
	touch /var/log/systemusage/memory.log;  
	echo "Fisierul memory.log a fost creat";  
fi  
printf "$(date +"%d %b %Y %T:") " >>/var/log/systemusage/memory.log  
#free - afiseaza datele referitoare la moemoria ram  
free | grep Mem | awk '{print  $2"\t"$3"\t"$4"\t"$5"\t"$6"\n"}' >>/var/log/systemusage/memory.log  
  
if [[ -f "var/log/systemusage/load.log" ]]  
then   
	echo "Fisierul load.log exista"  
else  
	touch /var/log/systemusage/load.log;  
	echo "Fisierul load.log a fost creat";  
fi  
printf "$(date +"%d %b %Y %T:") " >>/var/log/systemusage/load.log  
cat /proc/loadavg | cut -f1,2,3 -d" " >>/var/log/systemusage/load.log  
printf "\n"  
```
