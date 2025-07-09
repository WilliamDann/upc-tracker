import Resource from './Resource'

export default class Product extends Resource
{
    constructor()
    {
        super('places', { Name: "", Address: "", City: "", State: "" })
    }
}