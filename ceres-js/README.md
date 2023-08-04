# ceres-js    
    
This is a helper package to transfer data collected with a collector to a data DAO node. This package is written in JavaScript and can be intergrated into data collectors written in JavaScript.     
    
You can create data collectors in any language. An implementationof this data transport helper is planned in couple of other languages such as Rust, C++, and Swift.    
     
You can run this code both as a backend and in the browser. Pass in the `isBrowser` argument accordingly.   

The typical flow for transferring some data is as follows:
1. Call `cid` passing in your data to compute the content identifier for the data.   
2. Call `getDaoAdminPeerInfo` passing in the address of your data DAO admin.   
3. Call `verifyPeerInfo` to verify the authenticity of the peer info you received.    
4. Call `createCollectorMsg` to create the message that will be sent to the DAO node.    
4. Call `sendDataCollectorMsgToNode` to send the message to the DAO node. If the DAO node confirms the authenticity of the message the returned `error` will be `nil`.   
5. Call `transferFile` to transfer the data to the DAO node.   
    
## Methods    
```
setup = async (port, keyBytes) => {...}
```    
Creates and configures a new libp2p host. The host will open connections for our transport protocols on the given port.   
    
<br>
    
```
transferFile = async (targetMultiaddr, payload) => {...}
```   
Transfers the data from the given file to the given target address.   
    
<br>
    
```
sendDataCollectorMsgToNode = async (targetMultiaddr, msg) => {...}
```    
Sends a `DataCollectorMsg` to the given target address. If `error` is `nil`, the target accepted the message as authentic.   
    
<br>
    
```
cid = async (data) => {...}
```   
Generates the content identifier for the given data. Uses the Dag-Protobuf codec.   
    
<br>
    
```
createCollectorMsg = async (cid) => {...}
```    
A data collector message consists of a data CID and a signature over that CID with the Ethereum private key of the data collector. This function generates this message.   
    
<br>
    
```
getDaoAdminPeerInfo = async (baseUrl, address) => {...}
```    
Retrieves the `PeerInfo` for a given DAO admin address.
    
<br>
    

```
verifyPeerInfo = async(addr, mu, sig, isBrowser) => {...}
```
Verifies the authenticity of a `PeerInfo` object. The `PeerInfo` can then be used to communicate with that node.
