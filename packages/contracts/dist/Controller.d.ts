import { BaseContract, BigNumber, BigNumberish, BytesLike, CallOverrides, ContractTransaction, Overrides, PopulatedTransaction, Signer, utils } from "ethers";
import { FunctionFragment, Result } from "@ethersproject/abi";
import { Listener, Provider } from "@ethersproject/providers";
import { TypedEventFilter, TypedEvent, TypedListener, OnEvent } from "./common";
export interface ControllerInterface extends utils.Interface {
    functions: {
        "addAccessory(string)": FunctionFragment;
        "addBodyRles(uint8,bytes[])": FunctionFragment;
        "addItemComponent(uint8,string)": FunctionFragment;
        "batchSetItemRle(uint256[],bytes[])": FunctionFragment;
        "dao()": FunctionFragment;
        "initiator()": FunctionFragment;
        "maintainer()": FunctionFragment;
        "mintItem(address,uint8[5],uint8,uint256,bytes)": FunctionFragment;
        "mintItemBatch(address,uint8[],uint8[],uint256[],bytes)": FunctionFragment;
        "mintOGTo(uint256,address,string,bytes4,bytes4,bytes2,uint8[4],uint8[4],bytes2,bytes)": FunctionFragment;
        "mintTo(uint256,address,string,bytes4,bytes4,bytes2,uint8[4],uint8[4],bytes2,bytes)": FunctionFragment;
        "onERC1155BatchReceived(address,address,uint256[],uint256[],bytes)": FunctionFragment;
        "onERC1155Received(address,address,uint256,uint256,bytes)": FunctionFragment;
        "open(uint256,address,bytes)": FunctionFragment;
        "setDAO(address)": FunctionFragment;
        "setEnforcer(address)": FunctionFragment;
        "setInitiator(address)": FunctionFragment;
        "setItemRle(uint256,bytes,bytes)": FunctionFragment;
        "setMaintainer(address)": FunctionFragment;
        "setPalette(uint8,bytes4[])": FunctionFragment;
        "supportsInterface(bytes4)": FunctionFragment;
    };
    encodeFunctionData(functionFragment: "addAccessory", values: [string]): string;
    encodeFunctionData(functionFragment: "addBodyRles", values: [BigNumberish, BytesLike[]]): string;
    encodeFunctionData(functionFragment: "addItemComponent", values: [BigNumberish, string]): string;
    encodeFunctionData(functionFragment: "batchSetItemRle", values: [BigNumberish[], BytesLike[]]): string;
    encodeFunctionData(functionFragment: "dao", values?: undefined): string;
    encodeFunctionData(functionFragment: "initiator", values?: undefined): string;
    encodeFunctionData(functionFragment: "maintainer", values?: undefined): string;
    encodeFunctionData(functionFragment: "mintItem", values: [
        string,
        [
            BigNumberish,
            BigNumberish,
            BigNumberish,
            BigNumberish,
            BigNumberish
        ],
        BigNumberish,
        BigNumberish,
        BytesLike
    ]): string;
    encodeFunctionData(functionFragment: "mintItemBatch", values: [string, BigNumberish[], BigNumberish[], BigNumberish[], BytesLike]): string;
    encodeFunctionData(functionFragment: "mintOGTo", values: [
        BigNumberish,
        string,
        string,
        BytesLike,
        BytesLike,
        BytesLike,
        [
            BigNumberish,
            BigNumberish,
            BigNumberish,
            BigNumberish
        ],
        [
            BigNumberish,
            BigNumberish,
            BigNumberish,
            BigNumberish
        ],
        BytesLike,
        BytesLike
    ]): string;
    encodeFunctionData(functionFragment: "mintTo", values: [
        BigNumberish,
        string,
        string,
        BytesLike,
        BytesLike,
        BytesLike,
        [
            BigNumberish,
            BigNumberish,
            BigNumberish,
            BigNumberish
        ],
        [
            BigNumberish,
            BigNumberish,
            BigNumberish,
            BigNumberish
        ],
        BytesLike,
        BytesLike
    ]): string;
    encodeFunctionData(functionFragment: "onERC1155BatchReceived", values: [string, string, BigNumberish[], BigNumberish[], BytesLike]): string;
    encodeFunctionData(functionFragment: "onERC1155Received", values: [string, string, BigNumberish, BigNumberish, BytesLike]): string;
    encodeFunctionData(functionFragment: "open", values: [BigNumberish, string, BytesLike]): string;
    encodeFunctionData(functionFragment: "setDAO", values: [string]): string;
    encodeFunctionData(functionFragment: "setEnforcer", values: [string]): string;
    encodeFunctionData(functionFragment: "setInitiator", values: [string]): string;
    encodeFunctionData(functionFragment: "setItemRle", values: [BigNumberish, BytesLike, BytesLike]): string;
    encodeFunctionData(functionFragment: "setMaintainer", values: [string]): string;
    encodeFunctionData(functionFragment: "setPalette", values: [BigNumberish, BytesLike[]]): string;
    encodeFunctionData(functionFragment: "supportsInterface", values: [BytesLike]): string;
    decodeFunctionResult(functionFragment: "addAccessory", data: BytesLike): Result;
    decodeFunctionResult(functionFragment: "addBodyRles", data: BytesLike): Result;
    decodeFunctionResult(functionFragment: "addItemComponent", data: BytesLike): Result;
    decodeFunctionResult(functionFragment: "batchSetItemRle", data: BytesLike): Result;
    decodeFunctionResult(functionFragment: "dao", data: BytesLike): Result;
    decodeFunctionResult(functionFragment: "initiator", data: BytesLike): Result;
    decodeFunctionResult(functionFragment: "maintainer", data: BytesLike): Result;
    decodeFunctionResult(functionFragment: "mintItem", data: BytesLike): Result;
    decodeFunctionResult(functionFragment: "mintItemBatch", data: BytesLike): Result;
    decodeFunctionResult(functionFragment: "mintOGTo", data: BytesLike): Result;
    decodeFunctionResult(functionFragment: "mintTo", data: BytesLike): Result;
    decodeFunctionResult(functionFragment: "onERC1155BatchReceived", data: BytesLike): Result;
    decodeFunctionResult(functionFragment: "onERC1155Received", data: BytesLike): Result;
    decodeFunctionResult(functionFragment: "open", data: BytesLike): Result;
    decodeFunctionResult(functionFragment: "setDAO", data: BytesLike): Result;
    decodeFunctionResult(functionFragment: "setEnforcer", data: BytesLike): Result;
    decodeFunctionResult(functionFragment: "setInitiator", data: BytesLike): Result;
    decodeFunctionResult(functionFragment: "setItemRle", data: BytesLike): Result;
    decodeFunctionResult(functionFragment: "setMaintainer", data: BytesLike): Result;
    decodeFunctionResult(functionFragment: "setPalette", data: BytesLike): Result;
    decodeFunctionResult(functionFragment: "supportsInterface", data: BytesLike): Result;
    events: {};
}
export interface Controller extends BaseContract {
    connect(signerOrProvider: Signer | Provider | string): this;
    attach(addressOrName: string): this;
    deployed(): Promise<this>;
    interface: ControllerInterface;
    queryFilter<TEvent extends TypedEvent>(event: TypedEventFilter<TEvent>, fromBlockOrBlockhash?: string | number | undefined, toBlock?: string | number | undefined): Promise<Array<TEvent>>;
    listeners<TEvent extends TypedEvent>(eventFilter?: TypedEventFilter<TEvent>): Array<TypedListener<TEvent>>;
    listeners(eventName?: string): Array<Listener>;
    removeAllListeners<TEvent extends TypedEvent>(eventFilter: TypedEventFilter<TEvent>): this;
    removeAllListeners(eventName?: string): this;
    off: OnEvent<this>;
    on: OnEvent<this>;
    once: OnEvent<this>;
    removeListener: OnEvent<this>;
    functions: {
        addAccessory(component: string, overrides?: Overrides & {
            from?: string | Promise<string>;
        }): Promise<ContractTransaction>;
        addBodyRles(part: BigNumberish, _rles: BytesLike[], overrides?: Overrides & {
            from?: string | Promise<string>;
        }): Promise<ContractTransaction>;
        addItemComponent(componentType: BigNumberish, component: string, overrides?: Overrides & {
            from?: string | Promise<string>;
        }): Promise<ContractTransaction>;
        batchSetItemRle(ids: BigNumberish[], rles: BytesLike[], overrides?: Overrides & {
            from?: string | Promise<string>;
        }): Promise<ContractTransaction>;
        dao(overrides?: CallOverrides): Promise<[string]>;
        initiator(overrides?: CallOverrides): Promise<[string]>;
        maintainer(overrides?: CallOverrides): Promise<[string]>;
        mintItem(to: string, components_: [
            BigNumberish,
            BigNumberish,
            BigNumberish,
            BigNumberish,
            BigNumberish
        ], componentType: BigNumberish, amount: BigNumberish, data: BytesLike, overrides?: Overrides & {
            from?: string | Promise<string>;
        }): Promise<ContractTransaction>;
        mintItemBatch(to: string, components_: BigNumberish[], componentTypes: BigNumberish[], amounts: BigNumberish[], data: BytesLike, overrides?: Overrides & {
            from?: string | Promise<string>;
        }): Promise<ContractTransaction>;
        mintOGTo(dopeId: BigNumberish, to: string, name: string, color: BytesLike, background: BytesLike, options: BytesLike, viewbox: [BigNumberish, BigNumberish, BigNumberish, BigNumberish], body: [BigNumberish, BigNumberish, BigNumberish, BigNumberish], mask: BytesLike, data: BytesLike, overrides?: Overrides & {
            from?: string | Promise<string>;
        }): Promise<ContractTransaction>;
        mintTo(dopeId: BigNumberish, to: string, name: string, color: BytesLike, background: BytesLike, options: BytesLike, viewbox: [BigNumberish, BigNumberish, BigNumberish, BigNumberish], body: [BigNumberish, BigNumberish, BigNumberish, BigNumberish], mask: BytesLike, data: BytesLike, overrides?: Overrides & {
            from?: string | Promise<string>;
        }): Promise<ContractTransaction>;
        onERC1155BatchReceived(arg0: string, arg1: string, arg2: BigNumberish[], arg3: BigNumberish[], arg4: BytesLike, overrides?: CallOverrides): Promise<[string]>;
        onERC1155Received(operator: string, arg1: string, arg2: BigNumberish, arg3: BigNumberish, arg4: BytesLike, overrides?: CallOverrides): Promise<[string]>;
        open(dopeId: BigNumberish, to: string, data: BytesLike, overrides?: Overrides & {
            from?: string | Promise<string>;
        }): Promise<ContractTransaction>;
        setDAO(dao_: string, overrides?: Overrides & {
            from?: string | Promise<string>;
        }): Promise<ContractTransaction>;
        setEnforcer(enforcer_: string, overrides?: Overrides & {
            from?: string | Promise<string>;
        }): Promise<ContractTransaction>;
        setInitiator(initiator_: string, overrides?: Overrides & {
            from?: string | Promise<string>;
        }): Promise<ContractTransaction>;
        setItemRle(id: BigNumberish, male: BytesLike, female: BytesLike, overrides?: Overrides & {
            from?: string | Promise<string>;
        }): Promise<ContractTransaction>;
        setMaintainer(maintainer_: string, overrides?: Overrides & {
            from?: string | Promise<string>;
        }): Promise<ContractTransaction>;
        setPalette(id: BigNumberish, palette: BytesLike[], overrides?: Overrides & {
            from?: string | Promise<string>;
        }): Promise<ContractTransaction>;
        supportsInterface(interfaceId: BytesLike, overrides?: CallOverrides): Promise<[boolean]>;
    };
    addAccessory(component: string, overrides?: Overrides & {
        from?: string | Promise<string>;
    }): Promise<ContractTransaction>;
    addBodyRles(part: BigNumberish, _rles: BytesLike[], overrides?: Overrides & {
        from?: string | Promise<string>;
    }): Promise<ContractTransaction>;
    addItemComponent(componentType: BigNumberish, component: string, overrides?: Overrides & {
        from?: string | Promise<string>;
    }): Promise<ContractTransaction>;
    batchSetItemRle(ids: BigNumberish[], rles: BytesLike[], overrides?: Overrides & {
        from?: string | Promise<string>;
    }): Promise<ContractTransaction>;
    dao(overrides?: CallOverrides): Promise<string>;
    initiator(overrides?: CallOverrides): Promise<string>;
    maintainer(overrides?: CallOverrides): Promise<string>;
    mintItem(to: string, components_: [
        BigNumberish,
        BigNumberish,
        BigNumberish,
        BigNumberish,
        BigNumberish
    ], componentType: BigNumberish, amount: BigNumberish, data: BytesLike, overrides?: Overrides & {
        from?: string | Promise<string>;
    }): Promise<ContractTransaction>;
    mintItemBatch(to: string, components_: BigNumberish[], componentTypes: BigNumberish[], amounts: BigNumberish[], data: BytesLike, overrides?: Overrides & {
        from?: string | Promise<string>;
    }): Promise<ContractTransaction>;
    mintOGTo(dopeId: BigNumberish, to: string, name: string, color: BytesLike, background: BytesLike, options: BytesLike, viewbox: [BigNumberish, BigNumberish, BigNumberish, BigNumberish], body: [BigNumberish, BigNumberish, BigNumberish, BigNumberish], mask: BytesLike, data: BytesLike, overrides?: Overrides & {
        from?: string | Promise<string>;
    }): Promise<ContractTransaction>;
    mintTo(dopeId: BigNumberish, to: string, name: string, color: BytesLike, background: BytesLike, options: BytesLike, viewbox: [BigNumberish, BigNumberish, BigNumberish, BigNumberish], body: [BigNumberish, BigNumberish, BigNumberish, BigNumberish], mask: BytesLike, data: BytesLike, overrides?: Overrides & {
        from?: string | Promise<string>;
    }): Promise<ContractTransaction>;
    onERC1155BatchReceived(arg0: string, arg1: string, arg2: BigNumberish[], arg3: BigNumberish[], arg4: BytesLike, overrides?: CallOverrides): Promise<string>;
    onERC1155Received(operator: string, arg1: string, arg2: BigNumberish, arg3: BigNumberish, arg4: BytesLike, overrides?: CallOverrides): Promise<string>;
    open(dopeId: BigNumberish, to: string, data: BytesLike, overrides?: Overrides & {
        from?: string | Promise<string>;
    }): Promise<ContractTransaction>;
    setDAO(dao_: string, overrides?: Overrides & {
        from?: string | Promise<string>;
    }): Promise<ContractTransaction>;
    setEnforcer(enforcer_: string, overrides?: Overrides & {
        from?: string | Promise<string>;
    }): Promise<ContractTransaction>;
    setInitiator(initiator_: string, overrides?: Overrides & {
        from?: string | Promise<string>;
    }): Promise<ContractTransaction>;
    setItemRle(id: BigNumberish, male: BytesLike, female: BytesLike, overrides?: Overrides & {
        from?: string | Promise<string>;
    }): Promise<ContractTransaction>;
    setMaintainer(maintainer_: string, overrides?: Overrides & {
        from?: string | Promise<string>;
    }): Promise<ContractTransaction>;
    setPalette(id: BigNumberish, palette: BytesLike[], overrides?: Overrides & {
        from?: string | Promise<string>;
    }): Promise<ContractTransaction>;
    supportsInterface(interfaceId: BytesLike, overrides?: CallOverrides): Promise<boolean>;
    callStatic: {
        addAccessory(component: string, overrides?: CallOverrides): Promise<void>;
        addBodyRles(part: BigNumberish, _rles: BytesLike[], overrides?: CallOverrides): Promise<void>;
        addItemComponent(componentType: BigNumberish, component: string, overrides?: CallOverrides): Promise<void>;
        batchSetItemRle(ids: BigNumberish[], rles: BytesLike[], overrides?: CallOverrides): Promise<void>;
        dao(overrides?: CallOverrides): Promise<string>;
        initiator(overrides?: CallOverrides): Promise<string>;
        maintainer(overrides?: CallOverrides): Promise<string>;
        mintItem(to: string, components_: [
            BigNumberish,
            BigNumberish,
            BigNumberish,
            BigNumberish,
            BigNumberish
        ], componentType: BigNumberish, amount: BigNumberish, data: BytesLike, overrides?: CallOverrides): Promise<void>;
        mintItemBatch(to: string, components_: BigNumberish[], componentTypes: BigNumberish[], amounts: BigNumberish[], data: BytesLike, overrides?: CallOverrides): Promise<void>;
        mintOGTo(dopeId: BigNumberish, to: string, name: string, color: BytesLike, background: BytesLike, options: BytesLike, viewbox: [BigNumberish, BigNumberish, BigNumberish, BigNumberish], body: [BigNumberish, BigNumberish, BigNumberish, BigNumberish], mask: BytesLike, data: BytesLike, overrides?: CallOverrides): Promise<void>;
        mintTo(dopeId: BigNumberish, to: string, name: string, color: BytesLike, background: BytesLike, options: BytesLike, viewbox: [BigNumberish, BigNumberish, BigNumberish, BigNumberish], body: [BigNumberish, BigNumberish, BigNumberish, BigNumberish], mask: BytesLike, data: BytesLike, overrides?: CallOverrides): Promise<void>;
        onERC1155BatchReceived(arg0: string, arg1: string, arg2: BigNumberish[], arg3: BigNumberish[], arg4: BytesLike, overrides?: CallOverrides): Promise<string>;
        onERC1155Received(operator: string, arg1: string, arg2: BigNumberish, arg3: BigNumberish, arg4: BytesLike, overrides?: CallOverrides): Promise<string>;
        open(dopeId: BigNumberish, to: string, data: BytesLike, overrides?: CallOverrides): Promise<void>;
        setDAO(dao_: string, overrides?: CallOverrides): Promise<void>;
        setEnforcer(enforcer_: string, overrides?: CallOverrides): Promise<void>;
        setInitiator(initiator_: string, overrides?: CallOverrides): Promise<void>;
        setItemRle(id: BigNumberish, male: BytesLike, female: BytesLike, overrides?: CallOverrides): Promise<void>;
        setMaintainer(maintainer_: string, overrides?: CallOverrides): Promise<void>;
        setPalette(id: BigNumberish, palette: BytesLike[], overrides?: CallOverrides): Promise<void>;
        supportsInterface(interfaceId: BytesLike, overrides?: CallOverrides): Promise<boolean>;
    };
    filters: {};
    estimateGas: {
        addAccessory(component: string, overrides?: Overrides & {
            from?: string | Promise<string>;
        }): Promise<BigNumber>;
        addBodyRles(part: BigNumberish, _rles: BytesLike[], overrides?: Overrides & {
            from?: string | Promise<string>;
        }): Promise<BigNumber>;
        addItemComponent(componentType: BigNumberish, component: string, overrides?: Overrides & {
            from?: string | Promise<string>;
        }): Promise<BigNumber>;
        batchSetItemRle(ids: BigNumberish[], rles: BytesLike[], overrides?: Overrides & {
            from?: string | Promise<string>;
        }): Promise<BigNumber>;
        dao(overrides?: CallOverrides): Promise<BigNumber>;
        initiator(overrides?: CallOverrides): Promise<BigNumber>;
        maintainer(overrides?: CallOverrides): Promise<BigNumber>;
        mintItem(to: string, components_: [
            BigNumberish,
            BigNumberish,
            BigNumberish,
            BigNumberish,
            BigNumberish
        ], componentType: BigNumberish, amount: BigNumberish, data: BytesLike, overrides?: Overrides & {
            from?: string | Promise<string>;
        }): Promise<BigNumber>;
        mintItemBatch(to: string, components_: BigNumberish[], componentTypes: BigNumberish[], amounts: BigNumberish[], data: BytesLike, overrides?: Overrides & {
            from?: string | Promise<string>;
        }): Promise<BigNumber>;
        mintOGTo(dopeId: BigNumberish, to: string, name: string, color: BytesLike, background: BytesLike, options: BytesLike, viewbox: [BigNumberish, BigNumberish, BigNumberish, BigNumberish], body: [BigNumberish, BigNumberish, BigNumberish, BigNumberish], mask: BytesLike, data: BytesLike, overrides?: Overrides & {
            from?: string | Promise<string>;
        }): Promise<BigNumber>;
        mintTo(dopeId: BigNumberish, to: string, name: string, color: BytesLike, background: BytesLike, options: BytesLike, viewbox: [BigNumberish, BigNumberish, BigNumberish, BigNumberish], body: [BigNumberish, BigNumberish, BigNumberish, BigNumberish], mask: BytesLike, data: BytesLike, overrides?: Overrides & {
            from?: string | Promise<string>;
        }): Promise<BigNumber>;
        onERC1155BatchReceived(arg0: string, arg1: string, arg2: BigNumberish[], arg3: BigNumberish[], arg4: BytesLike, overrides?: CallOverrides): Promise<BigNumber>;
        onERC1155Received(operator: string, arg1: string, arg2: BigNumberish, arg3: BigNumberish, arg4: BytesLike, overrides?: CallOverrides): Promise<BigNumber>;
        open(dopeId: BigNumberish, to: string, data: BytesLike, overrides?: Overrides & {
            from?: string | Promise<string>;
        }): Promise<BigNumber>;
        setDAO(dao_: string, overrides?: Overrides & {
            from?: string | Promise<string>;
        }): Promise<BigNumber>;
        setEnforcer(enforcer_: string, overrides?: Overrides & {
            from?: string | Promise<string>;
        }): Promise<BigNumber>;
        setInitiator(initiator_: string, overrides?: Overrides & {
            from?: string | Promise<string>;
        }): Promise<BigNumber>;
        setItemRle(id: BigNumberish, male: BytesLike, female: BytesLike, overrides?: Overrides & {
            from?: string | Promise<string>;
        }): Promise<BigNumber>;
        setMaintainer(maintainer_: string, overrides?: Overrides & {
            from?: string | Promise<string>;
        }): Promise<BigNumber>;
        setPalette(id: BigNumberish, palette: BytesLike[], overrides?: Overrides & {
            from?: string | Promise<string>;
        }): Promise<BigNumber>;
        supportsInterface(interfaceId: BytesLike, overrides?: CallOverrides): Promise<BigNumber>;
    };
    populateTransaction: {
        addAccessory(component: string, overrides?: Overrides & {
            from?: string | Promise<string>;
        }): Promise<PopulatedTransaction>;
        addBodyRles(part: BigNumberish, _rles: BytesLike[], overrides?: Overrides & {
            from?: string | Promise<string>;
        }): Promise<PopulatedTransaction>;
        addItemComponent(componentType: BigNumberish, component: string, overrides?: Overrides & {
            from?: string | Promise<string>;
        }): Promise<PopulatedTransaction>;
        batchSetItemRle(ids: BigNumberish[], rles: BytesLike[], overrides?: Overrides & {
            from?: string | Promise<string>;
        }): Promise<PopulatedTransaction>;
        dao(overrides?: CallOverrides): Promise<PopulatedTransaction>;
        initiator(overrides?: CallOverrides): Promise<PopulatedTransaction>;
        maintainer(overrides?: CallOverrides): Promise<PopulatedTransaction>;
        mintItem(to: string, components_: [
            BigNumberish,
            BigNumberish,
            BigNumberish,
            BigNumberish,
            BigNumberish
        ], componentType: BigNumberish, amount: BigNumberish, data: BytesLike, overrides?: Overrides & {
            from?: string | Promise<string>;
        }): Promise<PopulatedTransaction>;
        mintItemBatch(to: string, components_: BigNumberish[], componentTypes: BigNumberish[], amounts: BigNumberish[], data: BytesLike, overrides?: Overrides & {
            from?: string | Promise<string>;
        }): Promise<PopulatedTransaction>;
        mintOGTo(dopeId: BigNumberish, to: string, name: string, color: BytesLike, background: BytesLike, options: BytesLike, viewbox: [BigNumberish, BigNumberish, BigNumberish, BigNumberish], body: [BigNumberish, BigNumberish, BigNumberish, BigNumberish], mask: BytesLike, data: BytesLike, overrides?: Overrides & {
            from?: string | Promise<string>;
        }): Promise<PopulatedTransaction>;
        mintTo(dopeId: BigNumberish, to: string, name: string, color: BytesLike, background: BytesLike, options: BytesLike, viewbox: [BigNumberish, BigNumberish, BigNumberish, BigNumberish], body: [BigNumberish, BigNumberish, BigNumberish, BigNumberish], mask: BytesLike, data: BytesLike, overrides?: Overrides & {
            from?: string | Promise<string>;
        }): Promise<PopulatedTransaction>;
        onERC1155BatchReceived(arg0: string, arg1: string, arg2: BigNumberish[], arg3: BigNumberish[], arg4: BytesLike, overrides?: CallOverrides): Promise<PopulatedTransaction>;
        onERC1155Received(operator: string, arg1: string, arg2: BigNumberish, arg3: BigNumberish, arg4: BytesLike, overrides?: CallOverrides): Promise<PopulatedTransaction>;
        open(dopeId: BigNumberish, to: string, data: BytesLike, overrides?: Overrides & {
            from?: string | Promise<string>;
        }): Promise<PopulatedTransaction>;
        setDAO(dao_: string, overrides?: Overrides & {
            from?: string | Promise<string>;
        }): Promise<PopulatedTransaction>;
        setEnforcer(enforcer_: string, overrides?: Overrides & {
            from?: string | Promise<string>;
        }): Promise<PopulatedTransaction>;
        setInitiator(initiator_: string, overrides?: Overrides & {
            from?: string | Promise<string>;
        }): Promise<PopulatedTransaction>;
        setItemRle(id: BigNumberish, male: BytesLike, female: BytesLike, overrides?: Overrides & {
            from?: string | Promise<string>;
        }): Promise<PopulatedTransaction>;
        setMaintainer(maintainer_: string, overrides?: Overrides & {
            from?: string | Promise<string>;
        }): Promise<PopulatedTransaction>;
        setPalette(id: BigNumberish, palette: BytesLike[], overrides?: Overrides & {
            from?: string | Promise<string>;
        }): Promise<PopulatedTransaction>;
        supportsInterface(interfaceId: BytesLike, overrides?: CallOverrides): Promise<PopulatedTransaction>;
    };
}
