
## Documentation
This page will provide information on why I implemented the challenge the way I did and will answer some bonus questions.


### Report
As a design pattern I chose the domain hex/ports and adapters.
I chose it because it's a perfect fit for modern application design. The only drawback is that it has a little overhead in code writing (takes time to create the different services).
However it allows for easy code injection for testing and allows for easy transition to other technologies i.e. instead of a kafka reader adapter it could be replaced with a aws sqs adapter without touching the actual service layer.


#### #input service
The input service is designed in a way that allows for continuous retrieval of incoming events from the adapter (in this case kafka). By using channels I could make use of a neat feature in go, which allows to read data of a channel in a loop, which will be in a "lock" and wait until there is a new event in the channel before it runs the code inside the loop. 
The data structure used here is a struct with just 2 values (ts and uid) since the other data is not needed for this challenge and allowed for faster development. Ideally you would parse all the payload and filter out what you need in the service.
Advantage of this implementation is that the parsing only needs to be done once but can be used for multiple sliding windows.
#### #counting service
The counting service is just counting the unique users in a given window. (It is my version of a sliding window algorithm, so might not be ideal) 
Sliding window allow to put an imaginary window on top of a data stream to access for example just a specific timeframe.  
The service has a ticker which will send a the counted amount for a finished window (defined time + 5 seconds) into a channel. (This part should be inside output service, but I did not have the time left to copy it over).
The windowSlides array holds a initial time of the first event that cam into the window and holds a usermap that is required to check wether a user already visited within the window.

#### #rest 
kafkaadapter holds reader and writer adapters for kafka
stdin is a simple filereader that will read a json file into an array and returns one element at a time to replicate a similar behaviour like kafka
stdout is a simple console logger
output service should have more logic (see counting service) to it but for now just forwards output to the adapter the service was initiated with
metric service outputs incoming events per second if there were any. (implementation not perfect)

### performance
between 9000 - 16000 fps running on Docker on M1 Apple Silicon
between 19000 - 30000 fps running directly on M1 Apple Silicon

If we assume that json is the limiting factor according to https://github.com/alecthomas/go_serialization_benchmarks using protobuf could be around 18x faster in unmarshalling. 
However for fast prototyping json is a good fit because it is humanly readable and fast enough for most things.

### error in counting
0.1% of messages arrive after 5 seconds
Additionally as I am writing this I noticed that if a message that arrived late would create a new sliding window that, depending on the timeframe, could include only 1 user, since it it uses the time of the event as a comparator, so the rate of error could be higher than 0,1% because of this. (would need to be adressed)

### scalability
 * use one service/application per time window
 * use loadbalancer and add a cache to store hits per user then use a different service to read out the values (not sure on this approach tho)
 
