import { IPAddress } from '../../interfaces/Index';

type Props = {
  address: IPAddress;
  deleteIP: () => void;
};

const ReportedIPAddress: React.FC<Props> = ({ address, deleteIP }) => {
  return (
    <div
      className={
        'relative flex flex-col break-words w-full mb-6 shadow-lg rounded bg-gray-700 text-white'
      }
    >
      <div className="px-4 py-3 mb-0 border-0 rounded-t">
        <div className="flex flex-wrap items-center">
          <div className="relative flex-1 flex-grow w-full max-w-full px-2">
            <h3 className={'font-semibold text-lg text-white'}>Servers</h3>
          </div>
        </div>
      </div>
      <div className="block w-full overflow-x-auto">
        <table className="items-center w-full bg-transparent border-collapse">
          <thead>
            <tr>
              <th
                className={
                  'px-6 align-middle border border-solid py-3 text-xs uppercase border-l-0 border-r-0 whitespace-nowrap font-semibold text-left bg-gray-600 text-gray-200 border-gray-500'
                }
              >
                Id
              </th>
              <th
                className={
                  'px-6 align-middle border border-solid py-3 text-xs uppercase border-l-0 border-r-0 whitespace-nowrap font-semibold text-left bg-gray-600 text-gray-200 border-gray-500'
                }
              >
                External IP
              </th>
              <th
                className={
                  'px-6 align-middle border border-solid py-3 text-xs uppercase border-l-0 border-r-0 whitespace-nowrap font-semibold text-left bg-gray-600 text-gray-200 border-gray-500'
                }
              >
                Created
              </th>
              <th
                className={
                  'px-6 align-middle border border-solid py-3 text-xs uppercase border-l-0 border-r-0 whitespace-nowrap font-semibold text-left bg-gray-600 text-gray-200 border-gray-500'
                }
              >
                Updated
              </th>
              <th
                className={
                  'px-6 align-middle border border-solid py-3 text-xs uppercase border-l-0 border-r-0 whitespace-nowrap font-semibold text-left bg-gray-600 text-gray-200 border-gray-500'
                }
              >
                Actions
              </th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td className="p-4 px-6 text-xs text-left align-middle border-t-0 border-l-0 border-r-0 whitespace-nowrap">
                <span className="text-white">{address.Id.S}</span>
              </td>
              <td className="p-4 px-6 text-xs align-middle border-t-0 border-l-0 border-r-0 whitespace-nowrap">
                <span className="text-white">{address.ExternalIP.S}</span>
              </td>
              <td className="p-4 px-6 text-xs align-middle border-t-0 border-l-0 border-r-0 whitespace-nowrap">
                {address.Created.S
                  ? new Date(parseInt(address.Created.S)).toLocaleString()
                  : 'Unknown'}
              </td>
              <td className="p-4 px-6 text-xs align-middle border-t-0 border-l-0 border-r-0 whitespace-nowrap">
                {address.Updated.S
                  ? new Date(parseInt(address.Updated.S)).toLocaleString()
                  : 'Unknown'}
              </td>
              <td className="p-4 px-6 text-xs align-middle border-t-0 border-l-0 border-r-0 whitespace-nowrap">
                <button
                  className="px-3 py-1 mb-1 mr-1 text-xs font-bold text-white uppercase transition-all duration-150 ease-linear bg-red-600 rounded outline-none active:bg-indigo-600 focus:outline-none"
                  type="button"
                  onClick={deleteIP}
                >
                  Delete
                </button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  );
};

export default ReportedIPAddress;
