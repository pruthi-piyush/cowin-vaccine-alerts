# cowin-vaccine-alerts
For giving lightning fast alerts on your local machine as soon as vaccine slot becomes available. <br />
Once you run this project, it will give you alerts on your Laptop/PC asap whenever your desired vaccine is available for your pincode.


## Run this project 

You can run this project on Windows, Linux and Mac.

Download the binary file by simply clicking **Download button** :- <br />
**MAC** users      - https://github.com/pruthi-piyush/cowin-vaccine-alerts/blob/main/bin/cowin-vaccine-alerts-mac <br />
**WINDOWS** users  - https://github.com/pruthi-piyush/cowin-vaccine-alerts/blob/main/bin/cowin-vaccine-alerts-windows.exe <br />
**LINUX** users   - https://github.com/pruthi-piyush/cowin-vaccine-alerts/blob/main/bin/cowin-vaccine-alerts-linux

Run this command in the folder where it gets downloaded ( **_You might need to give it permissions to execute_** ) <br />
**Example For Mac** -

```
./cowin-vaccine-alerts-mac --pincode 124001
```
Just replace the pincode value with the pincode for which you want to receive alerts.

## Filters

Only passing pincode is mandatory. Remaining parameters have following default values - <br />
```
age - 18 
dose - 1 
vaccine - "" (Alerts will be raised for all the vaccines)
```

Usage of ./cowin-vaccine-alerts-mac: <br />
```
  -age int 
    	Age - Enter 18 if you want 18+ alerts, 45 if you want 45+ alerts. 
    	  If you don't enter any value - default will be 18 (default 18)
 ```
 ```
  -dose int 
    	Dose No. for which you would like the alerts - Default is Dose 1 (default 1) 
  ```
  ```
  -pincode int 
    	Pin Code
  ```
  ```
  -vaccine string 
    	Vaccine ( Can be either COVISHIELD or COVAXIN ). 
    	  You can leave this empty if you want alerts for all the vaccines
   ```

## Cases

**Case 1** - _You want alerts of a specific vaccine_. <br />
         ```
         ./cowin-vaccine-alerts-mac --pincode 124001 --vaccine COVAXIN
         ```
         
**Case 2** - _You want alerts only for Dose 2._ <br />
         ```
         ./cowin-vaccine-alerts-mac --pincode 124001 --dose 2
         ```

**Case 3** - _You want alerts only for 18+ age_. <br />
         ```
         ./cowin-vaccine-alerts-mac --pincode 124001 --age 18
         ```
         
         
## Contribution

For contributing to this project, please write to pruthi.piyush@gmail.com.

