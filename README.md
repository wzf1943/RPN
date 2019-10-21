# RPN
RPN is a web service that can parse reverse polish notation string and then write result as JSON format.

## System Diagram
RPN will use master and slave mode to make sure whole system has high avaible feature. The distribuit system algorithm will use memberlist protocal. 

## API
###/parse
This api recieves single reverse polish notation string as input and return result as JSON format.

###/batch
This api recieves multiple reverse polish notation string in JSON format and return result as JSON format too.
