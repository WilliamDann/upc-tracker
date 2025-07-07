import Resource from './Resource'

export default class Product extends Resource
{
    constructor()
    {
        super('products', { Name: "", UPC: "", Desc: "" })
    }
}