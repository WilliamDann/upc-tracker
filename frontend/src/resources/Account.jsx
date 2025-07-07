import Resource from "./Resource";

export default class Account extends Resource {
    constructor()
    {
        super('account', { Email: "", Password: "", Name: "" })
    }
}