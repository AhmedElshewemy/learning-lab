package com.example.kafka;

import java.util.Properties;
import org.apache.kafka.clients.consumer.KafkaConsumer;

/**
 * Hello world!
 *
 */
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
    }
}
