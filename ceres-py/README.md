# ceres-py   
    
This is a helper package to transfer data collected with a collector to a data DAO node. This package is written in Python.    

You can create data collectors in any language. An implementation of this data transport helper is planned in a couple of other languages such as Rust, C++, and Swift.  

The typical flow for transferring some data is as follows:
1. Call `cid` passing in your data to compute the content identifier for the data.   
2. Call `get_dao_admin_peer_info` passing in the address of your data DAO admin.   
3. Call `verify_peer_info` to verify the authenticity of the peer info you received.    
4. Call `create_collector_msg` to create the message that will be sent to the DAO node.    
4. Call `send_data_collector_msg_to_node` to send the message to the DAO node. If the DAO node confirms the authenticity of the message the returned `error` will be `nil`.   
5. Call `transfer_file` to transfer the data to the DAO node.   
     
## Methods    
```
async def run(port, target, filepath)
```     
Creates and configures a new libp2p host. The host will open connections for our transport protocols on the given port and runs the whole process of sending data to a DAO node.   
    
<br>
    
```
async def transfer_file(host, target, data)
```   
Transfers the data from the given file to the given target address.   
    
<br>
    
```
async def send_data_collector_msg_to_node(host, target, msg)
```    
Sends a `DataCollectorMsg` to the given target address. If `error` is `nil`, the target accepted the message as authentic.   
    
<br>
    
```
def get_cid(data_bytes)
```    
Generates the content identifier for the given data. Uses the Dag-Protobuf codec.   
    
<br>
    
```
def create_collector_msg(cid, sig)
```   
A data collector message consists of a data CID and a signature over that CID with the Ethereum private key of the data collector. This function generates this message.   
    
<br>
    
```
def get_dao_admin_peer_info(baseUrl, address)
```    
Retrieves the `PeerInfo` for a given DAO admin address.    
    
<br>
    

