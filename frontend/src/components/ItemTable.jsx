import React from 'react';

const ItemTable = ({ items }) => {
  if (!items || items.length === 0) {
    return <p>No items to display.</p>;
  }

  // Get the headers from the keys of the first item
  const headers = Object.keys(items[0]);

  return (
    <div className="overflow-x-auto rounded-xl shadow-md">
      <table className="min-w-full table-auto border border-gray-300">
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
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default ItemTable;
