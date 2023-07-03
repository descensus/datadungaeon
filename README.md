
<p align="center">
  <img src="datadungaeon.png" />
</p>




# Data dung aeon..datadungaeon
MQTT datasucker (sensors etc) that feeds the subscribed topics into a database such as MySQL/MariaDB.

Useful if you just want your data in a plain stupid SQL database.

# Support

Aqara sensors
- temperature
- magnet breaker
- power plug

Other
- Pocsag-message from [pocsag pilemaster](https://github.com/descensus/pocsag-pilemaster).

# Example usage

Create a database.

Set your environment variables, append a sequential integer for multiple devices, then start the application.

Database tables will be created automatically.

```
export DATADUNGAEON_DB_HOST=
export DATADUNGAEON_DB_PORT=
export DATADUNGAEON_DB_USERNAME=
export DATADUNGAEON_DB_PASSWORD=
export DATADUNGAEON_DB_DBNAME=

export DATADUNGAEON_MQTT_CLIENTID=
export DATADUNGAEON_MQTT_HOST=
export DATADUNGAEON_MQTT_USERNAME=
export DATADUNGAEON_MQTT_PASSWORD=
export DATADUNGAEON_MQTT_PORT=


export AQARA_PLUG0="garage/plug"
export AQARA_PLUG1="mansion/plug"
export AQARA_PLUG2="nuclearbunker/plug"

export AQARA_MAGNET0="houseZigbee/farstu"
export AQARA_MAGNET1="officeZigbee/basement"

export AQARA_TEMPERATURE0="garage/office-temp"
export AQARA_TEMPERATURE1="mansion/vardagsrum-temp"

export POCSAG0="pocsag/pocsag01"
```
