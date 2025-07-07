import Resource from "./Resource";

export default class Account extends Resource {
    constructor()
    {
        super('accounts', { Email: "", Password: "", Name: "" })
    }
}