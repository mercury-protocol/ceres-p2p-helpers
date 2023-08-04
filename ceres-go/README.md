# ceres-go   
    
This is a helper package to transfer data collected with a collector to a data DAO node. This package is written in Go and can be intergrated into data collectors written in Go.    
    
You can create data collectors in any language. An implementationof this data transport helper is planned in couple of other languages such as Rust, C++, and Swift.  

The typical flow for transferring some data is as follows:
1. Call `Cid` passing in your data to compute the content identifier for the data.   
2. Call `GetDaoAdminPeerInfo` passing in the address of your data DAO admin.   
3. Call `VerifyPeerInfo` to verify the authenticity of the peer info you received.    
4. Call `CreateCollectorMsg` to create the message that will be sent to the DAO node.    
4. Call `SendDataCollectorMsgToNode` to send the message to the DAO node. If the DAO node confirms the authenticity of the message the returned `error` will be `nil`.   
5. Call `TransferFile` to transfer the data to the DAO node.   
    
## Methods     
```
func Init(port int) (*host.Host, error)
```     
Creates and configures a new libp2p host. The host will open connections for our transport protocols on the given port.   
    
<br>
```
func TransferFile(n host.Host, filename string, target ma.Multiaddr) error
```   
Transfers the data from the given file to the given target address.   
   
```
func SendDataCollectorMsgToNode(n host.Host, target ma.Multiaddr, msg DataCollectorMsg) error
```    
Sends a `DataCollectorMsg` to the given target address. If `error` is `nil`, the target accepted the message as authentic.   
   
```
func Cid(data []byte) (string, error)
```    
Generates the content identifier for the given data. Uses the Dag-Protobuf codec.   
   
```
func CreateCollectorMsg(cid string, k string) (DataCollectorMsg, error)
```   
A data collector message consists of a data CID and a signature over that CID with the Ethereum private key of the data collector. This function generates this message.   
   
```
func VerifyPeerInfo(p PeerInfo, clientUrl string) (bool, error)
```    
Verifies the authenticity of a `PeerInfo` object. The `PeerInfo` can then be used to communicate with that node.
