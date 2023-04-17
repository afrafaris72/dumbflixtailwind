import React, { useState } from 'react';
import { IoMdArrowDropdownCircle } from 'react-icons/io';

const AdminTransactions = () => {
  const [dropButton, setDropButton] = useState(false);
  return (
    <React.Fragment>
      <div className="px-20 bg-black py-10">
        <div className="container mx-auto py-20 h-[100vh]">
          <h2 className="font-bold text-white text-lg mb-5">Incoming Transaction</h2>
          <table className="table w-full">
            <thead>
              <tr>
                <th>No</th>
                <th>Users</th>
                <th>Bukti Transfer</th>
                <th>Remaining Active</th>
                <th>Status User</th>
                <th>Status Payment</th>
                <th>Action</th>
              </tr>
            </thead>
            <tbody>
              <tr>
                <td>1</td>
                <td>Radif Ganteng</td>
                <td>Bca.jpg</td>
                <td>26 /hari</td>
                <td className="text-green-500">Active</td>
                <td>Approve</td>
                <td className="relative">
                  <button className="pl-3" onClick={() => setDropButton(!dropButton)}>
                    <IoMdArrowDropdownCircle className="text-2xl text-sky-500" />
                  </button>
                  
                </td>
              </tr>
              <tr>
                <td>1</td>
                <td>Haris Rahman</td>
                <td>Bni.jpg</td>
                <td>26 /hari</td>
                <td className="text-green-500">Active</td>
                <td>Approve</td>
                <td className="relative">
                  <button className="pl-3" onClick={() => setDropButton(true)}>
                    <IoMdArrowDropdownCircle className="text-2xl text-sky-500" />
                  </button>
                  
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </React.Fragment>
  );
};

export default AdminTransactions;
