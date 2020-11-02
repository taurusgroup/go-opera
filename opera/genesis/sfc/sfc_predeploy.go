package sfc

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// GetContractBin is SFC contract genesis implementation bin code
// Has to be compiled with flag bin-runtime
func GetContractBin() []byte {
	return hexutil.MustDecode("0x6080604052600436106103345760003560e01c806376671808116101b0578063aa34eb45116100ec578063c7be95de11610095578063e66e1da21161006f578063e66e1da214610d69578063f2fde38b14610d93578063f9a1b8ab14610dc6578063fe4b84df14610df057610334565b8063c7be95de14610d06578063ca40edf114610d1b578063d9a7c1f914610d5457610334565b8063bc165f2d116100c6578063bc165f2d14610c4a578063bfb2c56314610c74578063c5f530af14610cf157610334565b8063aa34eb4514610b70578063b3341c4314610b9a578063b5d8962714610bdf57610334565b80638da5cb5b11610159578063a0dec17f11610133578063a0dec17f14610a23578063a5a470ad14610ace578063a694fc3a14610b3e578063a778651514610b5b57610334565b80638da5cb5b146109995780638f32d59b146109ca57806393095f95146109f357610334565b80638914d4c01161018a5780638914d4c0146104495780638b0e9f3f1461094b5780638cddb0151461096057610334565b806376671808146108825780637cacb1d614610897578063854873e1146108ac57610334565b80632e5f75ef1161027f5780635ccfe1e8116102285780636099ecb2116102025780636099ecb21461079e578063689dfca1146107d75780636f49866314610834578063715018a61461086d57610334565b80635ccfe1e8146107505780635e2308d2146104735780635fab23a81461078957610334565b80633fee10a8116102595780633fee10a8146104b657806350e4e9481461059357806354fd4d501461070657610334565b80632e5f75ef146104cb57806336e7e23b146104fb57806339b80c001461053157610334565b80631c3c60c8116102e15780632709275e116102bb5780632709275e146104735780632cedb097146104885780632d296a9b146104b657610334565b80631c3c60c8146104195780631d58179c146104495780632265f2841461045e57610334565b80630d4955e3116103125780630d4955e3146103da5780630d7b2609146103ef57806318160ddd1461040457610334565b80630135b1db14610339578063042c37b51461037e5780630962ef79146103b0575b600080fd5b34801561034557600080fd5b5061036c6004803603602081101561035c57600080fd5b50356001600160a01b0316610e1a565b60408051918252519081900360200190f35b34801561038a57600080fd5b506103ae600480360360408110156103a157600080fd5b5080359060200135610e2c565b005b3480156103bc57600080fd5b506103ae600480360360208110156103d357600080fd5b5035610ee8565b3480156103e657600080fd5b5061036c610fd7565b3480156103fb57600080fd5b5061036c610fe0565b34801561041057600080fd5b5061036c610fe7565b34801561042557600080fd5b506103ae6004803603604081101561043c57600080fd5b5080359060200135610fed565b34801561045557600080fd5b5061036c611085565b34801561046a57600080fd5b5061036c61108a565b34801561047f57600080fd5b5061036c61109c565b34801561049457600080fd5b5061049d6110b8565b6040805192835260208301919091528051918290030190f35b3480156104c257600080fd5b5061036c6110c2565b3480156104d757600080fd5b506103ae600480360360408110156104ee57600080fd5b50803590602001356110c9565b34801561050757600080fd5b506103ae6004803603606081101561051e57600080fd5b508035906020810135906040013561116b565b34801561053d57600080fd5b5061055b6004803603602081101561055457600080fd5b503561144a565b604080519788526020880196909652868601949094526060860192909252608085015260a084015260c0830152519081900360e00190f35b34801561059f57600080fd5b506103ae600480360360808110156105b657600080fd5b8101906020810181356401000000008111156105d157600080fd5b8201836020820111156105e357600080fd5b8035906020019184602083028401116401000000008311171561060557600080fd5b91939092909160208101903564010000000081111561062357600080fd5b82018360208201111561063557600080fd5b8035906020019184602083028401116401000000008311171561065757600080fd5b91939092909160208101903564010000000081111561067557600080fd5b82018360208201111561068757600080fd5b803590602001918460208302840111640100000000831117156106a957600080fd5b9193909290916020810190356401000000008111156106c757600080fd5b8201836020820111156106d957600080fd5b803590602001918460208302840111640100000000831117156106fb57600080fd5b50909250905061148c565b34801561071257600080fd5b5061071b6115b8565b604080517fffffff00000000000000000000000000000000000000000000000000000000009092168252519081900360200190f35b34801561075c57600080fd5b5061036c6004803603604081101561077357600080fd5b506001600160a01b0381351690602001356115dc565b34801561079557600080fd5b5061036c6115f9565b3480156107aa57600080fd5b5061036c600480360360408110156107c157600080fd5b506001600160a01b0381351690602001356115ff565b3480156107e357600080fd5b50610816600480360360608110156107fa57600080fd5b506001600160a01b03813516906020810135906040013561170f565b60408051938452602084019290925282820152519081900360600190f35b34801561084057600080fd5b5061036c6004803603604081101561085757600080fd5b506001600160a01b038135169060200135611741565b34801561087957600080fd5b506103ae61175e565b34801561088e57600080fd5b5061036c611819565b3480156108a357600080fd5b5061036c611822565b3480156108b857600080fd5b506108d6600480360360208110156108cf57600080fd5b5035611828565b6040805160208082528351818301528351919283929083019185019080838360005b838110156109105781810151838201526020016108f8565b50505050905090810190601f16801561093d5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561095757600080fd5b5061036c6118e1565b34801561096c57600080fd5b506103ae6004803603604081101561098357600080fd5b506001600160a01b0381351690602001356118e7565b3480156109a557600080fd5b506109ae611942565b604080516001600160a01b039092168252519081900360200190f35b3480156109d657600080fd5b506109df611951565b604080519115158252519081900360200190f35b3480156109ff57600080fd5b506103ae60048036036040811015610a1657600080fd5b5080359060200135611962565b348015610a2f57600080fd5b506103ae6004803603610100811015610a4757600080fd5b6001600160a01b0382351691602081013591810190606081016040820135640100000000811115610a7757600080fd5b820183602082011115610a8957600080fd5b80359060200191846001830284011164010000000083111715610aab57600080fd5b919350915080359060208101359060408101359060608101359060800135611be0565b6103ae60048036036020811015610ae457600080fd5b810190602081018135640100000000811115610aff57600080fd5b820183602082011115610b1157600080fd5b80359060200191846001830284011164010000000083111715610b3357600080fd5b509092509050611ca5565b6103ae60048036036020811015610b5457600080fd5b5035611d4e565b348015610b6757600080fd5b5061036c611d5c565b348015610b7c57600080fd5b506103ae60048036036020811015610b9357600080fd5b5035611d72565b348015610ba657600080fd5b506103ae60048036036080811015610bbd57600080fd5b506001600160a01b038135169060208101359060408101359060600135611e01565b348015610beb57600080fd5b50610c0960048036036020811015610c0257600080fd5b5035611e9b565b604080519788526020880196909652868601949094526060860192909252608085015260a08401526001600160a01b031660c0830152519081900360e00190f35b348015610c5657600080fd5b506108d660048036036020811015610c6d57600080fd5b5035611ee1565b348015610c8057600080fd5b506103ae60048036036020811015610c9757600080fd5b810190602081018135640100000000811115610cb257600080fd5b820183602082011115610cc457600080fd5b80359060200191846020830284011164010000000083111715610ce657600080fd5b509092509050611f67565b348015610cfd57600080fd5b5061036c611ff6565b348015610d1257600080fd5b5061036c612005565b348015610d2757600080fd5b5061036c60048036036040811015610d3e57600080fd5b506001600160a01b03813516906020013561200b565b348015610d6057600080fd5b5061036c612028565b348015610d7557600080fd5b506103ae60048036036020811015610d8c57600080fd5b503561202e565b348015610d9f57600080fd5b506103ae60048036036020811015610db657600080fd5b50356001600160a01b03166120c2565b348015610dd257600080fd5b506103ae60048036036020811015610de957600080fd5b5035612124565b348015610dfc57600080fd5b506103ae60048036036020811015610e1357600080fd5b50356121d6565b60686020526000908152604090205481565b3315610e7f576040805162461bcd60e51b815260206004820152600c60248201527f6e6f742063616c6c61626c650000000000000000000000000000000000000000604482015290519081900360640190fd5b80610ed1576040805162461bcd60e51b815260206004820152600c60248201527f77726f6e67207374617475730000000000000000000000000000000000000000604482015290519081900360640190fd5b610edb82826122de565b610ee482612124565b5050565b336000610ef582846115ff565b905080610f49576040805162461bcd60e51b815260206004820152600c60248201527f7a65726f20726577617264730000000000000000000000000000000000000000604482015290519081900360640190fd5b6001600160a01b0382166000908152606e60209081526040808320868452909152812055610f7683612345565b6001600160a01b0383166000818152607160209081526040808320888452909152808220939093559151909183156108fc02918491818181858888f19350505050158015610fc8573d6000803e3d6000fd5b50610fd2816123a2565b505050565b6301e133805b90565b6212750090565b60755481565b610ff5611951565b611046576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b604080518381526020810183905281517f95ae5488127de4bc98492f4487556e7af9f37eb4b6d5e94f6d849e03ff76cc7c929181900390910190a15050565b600390565b60006110946123db565b601002905090565b600060646110a86123db565b601e02816110b257fe5b04905090565b6072546073549091565b62093a8090565b6110d1611951565b611122576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b60738190556072829055604080518381526020810183905281517f702756a07c05d0bbfd06fc17b67951a5f4deb7bb6b088407e68a58969daf2a34929181900390910190a15050565b3361117681856123e7565b50600082116111cc576040805162461bcd60e51b815260206004820152600b60248201527f7a65726f20616d6f756e74000000000000000000000000000000000000000000604482015290519081900360640190fd5b6001600160a01b0381166000908152607060209081526040808320878452909152902054821115611244576040805162461bcd60e51b815260206004820152601060248201527f6e6f7420656e6f756768207374616b6500000000000000000000000000000000604482015290519081900360640190fd5b6001600160a01b0381166000908152606f602090815260408083208784528252808320868452909152902060020154156112c5576040805162461bcd60e51b815260206004820152601360248201527f7572494420616c72656164792065786973747300000000000000000000000000604482015290519081900360640190fd5b6001600160a01b0381166000908152607060209081526040808320878452825280832080548690039055606790915290206003015461130a908363ffffffff61247616565b600085815260676020526040902060030155606b5461132f908363ffffffff61247616565b606b5561133b846124bf565b8061134c575061134a846124fb565b155b6113875760405162461bcd60e51b81526004018080602001828103825260298152602001806137e26029913960400191505060405180910390fd5b611390846124fb565b61139f5761139f8460016122de565b6001600160a01b0381166000908152606f60209081526040808320878452825280832086845290915290206002018290556113d8611819565b6001600160a01b0382166000908152606f60209081526040808320888452825280832087845290915290205561140c61252d565b6001600160a01b0382166000908152606f60209081526040808320888452825280832087845290915290206001015561144484612124565b50505050565b607660205280600052604060002060009150905080600701549080600801549080600901549080600a01549080600b01549080600c01549080600d0154905087565b33156114df576040805162461bcd60e51b815260206004820152600c60248201527f6e6f742063616c6c61626c650000000000000000000000000000000000000000604482015290519081900360640190fd5b6115ae88888080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525050604080516020808c0282810182019093528b82529093508b92508a91829185019084908082843760009201919091525050604080516020808b0282810182019093528a82529093508a92508991829185019084908082843760009201919091525050604080516020808a0282810182019093528982529093508992508891829185019084908082843760009201919091525061253192505050565b5050505050505050565b7f323032000000000000000000000000000000000000000000000000000000000090565b607160209081526000928352604080842090915290825290205481565b606c5481565b6001600160a01b0382166000908152607160209081526040808320848452825280832054808452607680845282852086865260010190935290832054909183908161164987612345565b8152602001908152602001600020600101600086815260200190815260200160002054905060006116cb61167b6123db565b6001600160a01b03891660009081526070602090815260408083208b84529091529020546116bf906116b3868863ffffffff61247616565b9063ffffffff6125ee16565b9063ffffffff61264716565b6001600160a01b0388166000908152606e602090815260408083208a8452909152902054909150611702908263ffffffff61268916565b9450505050505b92915050565b606f60209081526000938452604080852082529284528284209052825290208054600182015460029092015490919083565b606e60209081526000928352604080842090915290825290205481565b611766611951565b6117b7576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b6033546040516000916001600160a01b0316907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3603380547fffffffffffffffffffffffff0000000000000000000000000000000000000000169055565b60665460010190565b60665481565b60696020908152600091825260409182902080548351601f60027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff610100600186161502019093169290920491820184900484028101840190945280845290918301828280156118d95780601f106118ae576101008083540402835291602001916118d9565b820191906000526020600020905b8154815290600101906020018083116118bc57829003601f168201915b505050505081565b606b5481565b6118f182826123e7565b610ee4576040805162461bcd60e51b815260206004820152601060248201527f6e6f7468696e6720746f20737461736800000000000000000000000000000000604482015290519081900360640190fd5b6033546001600160a01b031690565b6033546001600160a01b0316331490565b3361196b613643565b506001600160a01b0381166000908152606f6020908152604080832086845282528083208584528252918290208251606081018452815480825260018301549382019390935260029091015492810192909252611a0f576040805162461bcd60e51b815260206004820152601560248201527f7265717565737420646f65736e27742065786973740000000000000000000000604482015290519081900360640190fd5b60208082015182516000878152606790935260409092206001015490919015801590611a4b575060008681526067602052604090206001015482115b15611a6c575050600084815260676020526040902060018101546002909101545b611a746110c2565b8201611a7e61252d565b1015611ad1576040805162461bcd60e51b815260206004820152601660248201527f6e6f7420656e6f7567682074696d652070617373656400000000000000000000604482015290519081900360640190fd5b611ad9611085565b8101611ae3611819565b1015611b36576040805162461bcd60e51b815260206004820152601860248201527f6e6f7420656e6f7567682065706f636873207061737365640000000000000000604482015290519081900360640190fd5b6001600160a01b0384166000908152606f602090815260408083208984528252808320888452825280832060028101805485835560019092018590558490558984526067909252822054909190608016151580611bc9576040516001600160a01b0388169084156108fc029085906000818181858888f19350505050158015611bc3573d6000803e3d6000fd5b50611bcd565b8291505b50606c8054909101905550505050505050565b600054610100900460ff1680611bf95750611bf96126e3565b80611c07575060005460ff16155b611c425760405162461bcd60e51b815260040180806020018281038252602e8152602001806137b4602e913960400191505060405180910390fd5b611c8a898989898080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508b92508a915089905088886126e9565b606a54881115611c9a57606a8890555b505050505050505050565b611cad611ff6565b341015611d01576040805162461bcd60e51b815260206004820152601760248201527f696e73756666696369656e742073656c662d7374616b65000000000000000000604482015290519081900360640190fd5b611d413383838080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061288292505050565b610ee433606a54346128ad565b611d593382346128ad565b50565b60006064611d686123db565b600f02816110b257fe5b611d7a611951565b611dcb576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b6040805182815290517f35feeeac858525cae277d98c1c4792d0550aeab30f107addc09d8d5279faa53f9181900360200190a150565b600054610100900460ff1680611e1a5750611e1a6126e3565b80611e28575060005460ff16155b611e635760405162461bcd60e51b815260040180806020018281038252602e8152602001806137b4602e913960400191505060405180910390fd5b611e6e848484612915565b6001600160a01b039093166000908152606e60209081526040808320948352939052919091209190915550565b606760205260009081526040902080546001820154600283015460038401546004850154600586015460069096015494959394929391929091906001600160a01b031687565b606d6020908152600091825260409182902080548351601f60027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff610100600186161502019093169290920491820184900484028101840190945280845290918301828280156118d95780601f106118ae576101008083540402835291602001916118d9565b3315611fba576040805162461bcd60e51b815260206004820152600c60248201527f6e6f742063616c6c61626c650000000000000000000000000000000000000000604482015290519081900360640190fd5b610ee4828280806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250612ab892505050565b6a02a055184a310c1260000090565b606a5481565b607060209081526000928352604080842090915290825290205481565b60745481565b612036611951565b612087576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b60748190556040805182815290517f8cd9dae1bbea2bc8a5e80ffce2c224727a25925130a03ae100619a8861ae23969181900360200190a150565b6120ca611951565b61211b576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b611d5981612b7b565b61212d81612c34565b61217e576040805162461bcd60e51b815260206004820152601760248201527f76616c696461746f7220646f65736e2774206578697374000000000000000000604482015290519081900360640190fd5b6000818152606760205260409020600381015490541561219c575060005b60408051828152905183917f73f1a7b46df3dcd64f307585eeca93bc7f50aeec84f79927698d6ca79564cb81919081900360200190a25050565b600054610100900460ff16806121ef57506121ef6126e3565b806121fd575060005460ff16155b6122385760405162461bcd60e51b815260040180806020018281038252602e8152602001806137b4602e913960400191505060405180910390fd5b600054610100900460ff1615801561229e57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff909116610100171660011790555b6122a733612c4b565b60668290558015610ee457600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff1690555050565b600082815260676020526040902054811115610ee457600082815260676020526040902081815560020154610ee45761231561252d565b60008381526067602052604090206001015561232f611819565b6000838152606760205260409020600201555050565b60008181526067602052604081206002015415612398576000828152606760205260409020600201546066541015612380575060665461239d565b5060008181526067602052604090206002015461239d565b506066545b919050565b60408051828152905130917f0aafa333cb618aab22abe53de00db4ab4f9cedec55ea67e86d304d0947af4a22919081900360200190a250565b670de0b6b3a764000090565b6000806123f484846115ff565b90506123ff83612345565b6001600160a01b0385166000818152607160209081526040808320888452825280832094909455918152606e82528281208682529091522054811415612449576000915050611709565b6001600160a01b0384166000908152606e6020908152604080832086845290915290205550600192915050565b60006124b883836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250612dac565b9392505050565b60006124e06124cc6123db565b6116bf6124d761108a565b6116b3866124fb565b60008381526067602052604090206003015411159050919050565b6000818152606760209081526040808320600601546001600160a01b0316835260708252808320938352929052205490565b4290565b60006076600061253f611819565b8152602001908152602001600020905060608160060180548060200260200160405190810160405280929190818152602001828054801561259f57602002820191906000526020600020905b81548152602001906001019080831161258b575b505050505090506125b282828888612e43565b6125be82828686612f57565b6125c6611819565b6066556125d161252d565b600783015550607454600b820155607554600d9091015550505050565b6000826125fd57506000611709565b8282028284828161260a57fe5b04146124b85760405162461bcd60e51b81526004018080602001828103825260218152602001806137936021913960400191505060405180910390fd5b60006124b883836040518060400160405280601a81526020017f536166654d6174683a206469766973696f6e206279207a65726f000000000000815250613521565b6000828201838110156124b8576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b303b1590565b6001600160a01b03881660009081526068602052604090205415612754576040805162461bcd60e51b815260206004820152601860248201527f76616c696461746f7220616c7265616479206578697374730000000000000000604482015290519081900360640190fd5b6001600160a01b03881660008181526068602090815260408083208b90558a8352606782528083208981556004810189905560058101889055600181018690556002810187905560060180547fffffffffffffffffffffffff0000000000000000000000000000000000000000169094179093556069815291902087516127dd92890190613664565b50867fb4f79f036aba7d0ead3f22c57d59dc58e9272d32d274c1476fe523e7ee6f027f876040518080602001828103825283818151815260200191508051906020019080838360005b8381101561283e578181015183820152602001612826565b50505050905090810190601f16801561286b5780820380516001836020036101000a031916815260200191505b509250505060405180910390a25050505050505050565b606a805460010190819055610fd2838284600061289d611819565b6128a561252d565b6000806126e9565b6000828152606760205260409020541561290e576040805162461bcd60e51b815260206004820152601660248201527f76616c696461746f722069736e27742061637469766500000000000000000000604482015290519081900360640190fd5b610fd28383835b6000811161296a576040805162461bcd60e51b815260206004820152600b60248201527f7a65726f20616d6f756e74000000000000000000000000000000000000000000604482015290519081900360640190fd5b61297382612c34565b6129c4576040805162461bcd60e51b815260206004820152601760248201527f76616c696461746f7220646f65736e2774206578697374000000000000000000604482015290519081900360640190fd5b6129ce83836123e7565b506001600160a01b0383166000908152607060209081526040808320858452909152902054612a03908263ffffffff61268916565b6001600160a01b0384166000908152607060209081526040808320868452825280832093909355606790522060030154612a43908263ffffffff61268916565b600083815260676020526040902060030155606b54612a68908263ffffffff61268916565b606b55612a74826124bf565b612aaf5760405162461bcd60e51b81526004018080602001828103825260298152602001806137e26029913960400191505060405180910390fd5b610fd282612124565b600060766000612ac6611819565b8152602001908152602001600020905060008090505b8251811015612b6557600060676000858481518110612af757fe5b6020026020010151815260200190815260200160002060030154905080836000016000868581518110612b2657fe5b6020026020010151815260200190815260200160002081905550612b578184600c015461268990919063ffffffff16565b600c84015550600101612adc565b508151610fd290600683019060208501906136e2565b6001600160a01b038116612bc05760405162461bcd60e51b815260040180806020018281038252602681526020018061376d6026913960400191505060405180910390fd5b6033546040516001600160a01b038084169216907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3603380547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b600090815260676020526040902060050154151590565b600054610100900460ff1680612c645750612c646126e3565b80612c72575060005460ff16155b612cad5760405162461bcd60e51b815260040180806020018281038252602e8152602001806137b4602e913960400191505060405180910390fd5b600054610100900460ff16158015612d1357600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff909116610100171660011790555b603380547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0384811691909117918290556040519116906000907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a38015610ee457600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff1690555050565b60008184841115612e3b5760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015612e00578181015183820152602001612de8565b50505050905090810190601f168015612e2d5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b60005b8351811015612f5057607254828281518110612e5e57fe5b6020026020010151118015612e885750607354838281518110612e7d57fe5b602002602001015110155b15612ec757612eab848281518110612e9c57fe5b602002602001015160086122de565b612ec7848281518110612eba57fe5b6020026020010151612124565b828181518110612ed357fe5b6020026020010151856004016000868481518110612eed57fe5b6020026020010151815260200190815260200160002081905550818181518110612f1357fe5b6020026020010151856005016000868481518110612f2d57fe5b602090810291909101810151825281019190915260400160002055600101612e46565b5050505050565b612f5f61371c565b6040518060c001604052808551604051908082528060200260200182016040528015612f95578160200160208202803883390190505b508152602001600081526020018551604051908082528060200260200182016040528015612fcd578160200160208202803883390190505b5081526020016000815260200160008152602001600081525090506000607660006130076001612ffb611819565b9063ffffffff61247616565b8152602081019190915260400160002060016080840152600781015490915061302e61252d565b111561304857806007015461304161252d565b0360808301525b60005b855181101561311e5761309583608001516116bf87848151811061306b57fe5b602002602001015187858151811061307f57fe5b60200260200101516125ee90919063ffffffff16565b836040015182815181106130a557fe5b6020026020010181815250506130df836040015182815181106130c457fe5b6020026020010151846060015161268990919063ffffffff16565b60608401528351613111908590839081106130f657fe5b60200260200101518460a0015161268990919063ffffffff16565b60a084015260010161304b565b5060005b85518110156131f5576131a083608001516116bf87848151811061314257fe5b60200260200101516116b387608001516116bf8b888151811061316157fe5b60200260200101518e60000160008f8b8151811061317b57fe5b60200260200101518152602001908152602001600020546125ee90919063ffffffff16565b83518051839081106131ae57fe5b6020026020010181815250506131e8836000015182815181106131cd57fe5b6020026020010151846020015161268990919063ffffffff16565b6020840152600101613122565b5060005b85518110156134f957600061323184608001516074548660000151858151811061321f57fe5b60200260200101518760200151613586565b905061326d6132608560a001518660400151858151811061324e57fe5b602002602001015187606001516135c9565b829063ffffffff61268916565b905060006132828261327d611d5c565b613626565b90506000606e6000606760008c888151811061329a57fe5b6020026020010151815260200190815260200160002060060160009054906101000a90046001600160a01b03166001600160a01b03166001600160a01b0316815260200190815260200160002060008a86815181106132f557fe5b60200260200101518152602001908152602001600020549050613321828261268990919063ffffffff16565b905080606e6000606760008d898151811061333857fe5b6020026020010151815260200190815260200160002060060160009054906101000a90046001600160a01b03166001600160a01b03166001600160a01b0316815260200190815260200160002060008b878151811061339357fe5b60200260200101518152602001908152602001600020819055506000828403905060006133d4606b546116bf6133c76123db565b859063ffffffff6125ee16565b9050613411818860010160008e8a815181106133ec57fe5b602002602001015181526020019081526020016000205461268990919063ffffffff16565b8c60010160008d898151811061342357fe5b602002602001015181526020019081526020016000208190555061346689878151811061344c57fe5b60200260200101518860030160008e8a815181106133ec57fe5b8c60030160008d898151811061347857fe5b60200260200101518152602001908152602001600020819055506134bb8a87815181106134a157fe5b60200260200101518860020160008e8a815181106133ec57fe5b8c60020160008d89815181106134cd57fe5b6020026020010151815260200190815260200160002081905550505050505080806001019150506131f9565b505060a081015160088601556020810151600986015560600151600a90940193909355505050565b600081836135705760405162461bcd60e51b8152602060048201818152835160248401528351909283926044909101919085019080838360008315612e00578181015183820152602001612de8565b50600083858161357c57fe5b0495945050505050565b600082613595575060006135c1565b60006135a7868663ffffffff6125ee16565b90506135bd836116bf838763ffffffff6125ee16565b9150505b949350505050565b6000826135d8575060006124b8565b60006135ee836116bf878763ffffffff6125ee16565b905061361d6135fb6123db565b6116bf61360661109c565b61360e6123db565b8591900363ffffffff6125ee16565b95945050505050565b60006124b86136336123db565b6116bf858563ffffffff6125ee16565b60405180606001604052806000815260200160008152602001600081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106136a557805160ff19168380011785556136d2565b828001600101855582156136d2579182015b828111156136d25782518255916020019190600101906136b7565b506136de929150613752565b5090565b8280548282559060005260206000209081019282156136d257916020028201828111156136d25782518255916020019190600101906136b7565b6040518060c001604052806060815260200160008152602001606081526020016000815260200160008152602001600081525090565b610fdd91905b808211156136de576000815560010161375856fe4f776e61626c653a206e6577206f776e657220697320746865207a65726f2061646472657373536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f77436f6e747261637420696e7374616e63652068617320616c7265616479206265656e20696e697469616c697a656476616c696461746f7227732064656c65676174696f6e73206c696d6974206973206578636565646564a265627a7a72315820bc48a8a74699fd487f0d0ad3a07c2ab4ba39ec20eea651e3adf7a70b2219e0e264736f6c634300050c0032")
}

// ContractAddress is the SFC proxy contract address
var ContractAddress = common.HexToAddress("0xfc00face00000000000000000000000000000000")

// the SFC contract first implementation address
var ContractAddressV1 = common.HexToAddress("0xfc00beef00000000000000000000000000000101")
