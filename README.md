# sjmq

KubeMQ client 實作、定義通用 event models 及 sender & receiver 抽象介面

Implement base on:

1. Outbox https://microservices.io/patterns/data/transactional-outbox.html
1. Polling-publisher https://microservices.io/patterns/data/polling-publisher.html

# Notifier

Notify the **polling publisher** there are new messages in outbox.

1. Wait for sender call.
1. Enable.
1. Send notification to **polling publisher**.
1. If get error response, retry **2.** till success.
1. Disable. (if there's any new message added during enable time, it won't cause extra notification.)

# Sender

Save message in outbox, and call notifier (Implement base on **one context per request**)

1. New **sender** per context.
2. store the message in outbox. (if not in transaction, call notifier)
3. call notifier after transaction commit.

# Receiver

Save message in inbox, and call the handler.

1. get message from MQ.
1. save to inbox. (should send an ACK to MQ)
1. call service to handle it.
1. service pick up the top message from inbox and process the message.
    1. if not success, wait a second, then try again.
    1. after 3 times failure, put message in the hospital. // should have another flow to fix these message
1. if success, delete the message from inbox.
1. go back to step 3. till the inbox is empty.