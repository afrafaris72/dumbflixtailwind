import React from 'react';

const Payment = () => {
  return (
    <React.Fragment>
      <div className="px-8 py-24 container mx-auto h-[100vh]">
        <h2 className="text-center text-3xl font-bold">Choose You're plan!</h2>
        <div className="flex gap-x-20 px-20 mt-10">
          
          {/* <div className="w-1/3 flex flex-col justify-between bg-zinc-800 shadow-md rounded-md p-3">
            <h3 className="font-bold text-md mb-5">King Box</h3>
            <p className="font-bold text-red-700 text-2xl mb-3">Rp 1.000.000 </p>
            <p className="text-justify">GOD plan, same as before but you'll be GOD and never know when your subscriptions over and when we take you're money.</p>
            <button className="bg-red-700 text-white mt-5 py-2 rounded-md">Buy plan</button>
          </div> */}
        </div>
      </div>
    </React.Fragment>
  );
};

export default Payment;
