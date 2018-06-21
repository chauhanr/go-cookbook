# Cryptography

Cryptography is a practice of securing communication even when a third party can view those communications.
There are two way (symmetric and asymmetric) methods, as well as one way hashing algorithms that can be used
to secure the message.

## Hashing
hashing is a method to convert a variable length message into a unique fixed length alphanumeric string. Algorithms are
MD5, SHA1, SHA256 etc. The process is non reversible which means once the message is hashed we cannot get the message back
at the other end.

## Encryption
Encryption on the other hand is different from hashing in the sense that the message can be retrieved using a key. There are two ways

* Symmetric Encryption - here a shared key is used at both ends to encrypt and decrypt the message. AES is an example of symmetric encryption
* Asymmetric Encryption - here a public and private key pair are used to secure communication. RSA, DSA and ECDSA are algorithm used in this form of encryption. Some applications that use this are:

    * PGP (pretty good privacy)
    * SSH (secure shell)
    * SSL (secure socket layer) this is an older version of communication between machines. The current version is called TLS (transport layer security)

#### Public and Private Key Pair
In order to communicate using the asymmetric scheme we need to have a public/private key and that is done using the RSA or ECDSA algorithms.
The x509 package in the Go crypto package can be used to generate the key pair. We can also use the open ssl software to do the same.

Once the keys are prepared we can use them to digitally sign the messages and communicate between the parties.

#### Asymmetric encryption and internet communication.
With the RSA algorithm you can only encrypt message that are equal to or less than the length of the key which is 2048 bits or 256 bytes. Due to this reason
the RSA algorithm cannot be used to encrypt entire messages which use the TLS scheme. To over come this drawback TLS/SSL use a combination of symmetric and asymmetric methods

1. The initial handshake between the two parties happens with the asymmetric method and a public and private keys are used to do the handshake. A symmetric key is transported to the client as part of the handshake.
2. Once the symmetric key has been shared using the RSA method the rest of the communication happens using AES scheme.

The public and private keys are stored in encoded files called PEM files.


### TLS server and clients
We can easily create TLS server and clients using the go packages

```
 func main(){
     certFile, privKeyFile, hostString := checkArgs() // capture the values
     serverCert, err := tls.LoadX509KeyPair(certFile, privKeyFile)
     if err != nil{
        ..
     }

     config := &tls.Config{
        Certificates : []tls.Certificate{serverCert},
        // loading the keys and certifactes the server will use
     }

     listener, err := tls.Listen("tcp", hostString, config)
     if err != nil{
         ...
     }
     defer listener.Close()

     for {
        clientConnection, err :=  listener.Accept()
        if err != nil {
            ..
        }
        go handleConnection(clientConnection)
     }
 }

func handleConnection(clientConn net.Conn){
    defer clientConn.Close()
    socketReader := bufio.NewReader(clientConn)

   for{
         // handle the message and client connection using socket reader
         message, err := socketReader.ReadString("\n")
         // work with message
         ... 
   }
}

```
