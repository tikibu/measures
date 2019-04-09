Measure storage
=================

This lib implements a simple in-memory measure storage which is optionally synced every N seconds with a persistent storage.   

It's thread-safe and as simple as possible. 

How to use
----------

To initialize:
 
	measures.Measures = measures.CreateMeasureStore(measures.NewRedisIncrementer("localhost", 6379),
		time.Second*10)
		
By default, an in-memory measure storage is used. 

To write:

		measures.Measures.Inc("some_measure_1234234", 1.0)		
		
		
To read:
        
        	val, err := measures.Measures.Get("some_measure_1234234")
        	if err != nil {
        	    log.Fatal("error has occured", err)
        	}
        	
        	
It supports lazy loading, meaning that after restart, the measure will only be retrieved after it was first accessed. 
		
		
