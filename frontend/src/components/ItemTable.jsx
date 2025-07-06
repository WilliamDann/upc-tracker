import { Link } from 'react-router-dom'

const ItemTable = ({ items }) => {
  if (!items || items.length === 0) {
    return <p>No items to display.</p>;
  }

  // Get the headers from the keys of the first item
  const headers = Object.keys(items[0]);

  return (
    <div className="w-100">
      <table className="w-100 table-auto border border-gray-300">
        <thead className="bg-gray-100">
          <tr>
            {headers.map((header) => (
              <th key={header} className="px-4 py-2 text-left font-medium text-gray-700 border-b">
                {header}
              </th>
            ))}
          </tr>
        </thead>
        <tbody>
          {items.map((item, index) => (
            <tr key={index} className="odd:bg-white even:bg-gray-50">
              {headers.map((header) => (
                <td key={header} className="px-4 py-2 border-b">
                  {item[header]}
                </td>
              ))}
              <td>
              <Link to={`view/${item.id}`} className="btn btn-primary m-1 p-1">View</Link>
              <Link to={`edit/${item.id}`} className="btn btn-primary m-1 p-1">Edit</Link>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default ItemTable;
