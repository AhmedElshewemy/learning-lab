package com.example.kafka;

import java.time.Duration;
import java.util.Collections;
import java.util.Properties;
import java.util.Random;

import org.apache.kafka.clients.consumer.ConsumerRecord;
import org.apache.kafka.clients.consumer.ConsumerRecords;
import org.apache.kafka.clients.consumer.KafkaConsumer;


public class App 
{
    private static String kafkaHost = "localhost:9092";
    private static String topic = "fancy-topic";
    private static String groupName = "fancy-group";
    public static void main( String[] args )
    {
        Properties properties = new Properties();
        properties.put("bootstrap.servers", kafkaHost);
        properties.put("group.id", groupName);
        properties.put("key.deserializer", "org.apache.kafka.common.serialization.StringDeserializer");
        properties.put("value.deserializer", "org.apache.kafka.common.serialization.StringDeserializer");
        KafkaConsumer<String, String> consumer = new KafkaConsumer<>(properties);
        consumer.subscribe(Collections.singletonList(topic));

        try{
            System.out.printf("Subscribed to topic: %s\n", topic);
            while(true){
                ConsumerRecords<String, String> records = consumer.poll(Duration.ofSeconds(10));
                for(ConsumerRecord<String, String> record : records){
                    System.out.printf("Consumed record with key %s and value %s, from partition %d, offset %d\n", 
                    record.key(),record.value(), record.partition(), record.offset());
                    Thread.sleep(new Random().nextInt(5000));
                
                }
            }
        }catch(Exception e){
            System.out.println("An error consuming messages: " + e);
            e.printStackTrace();
        }finally{
            consumer.close();
            System.out.println("Consumer closed");
        }
}
}