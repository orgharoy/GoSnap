import React, { useEffect, useState }  from 'react'
import '../customStyles/login.css'

const Auth = () => {
    const [pageTitle, setPageTitle] = useState("GoSnap | Login");
    
    useEffect(() => {
        document.title = pageTitle;
    }, [pageTitle]);

    return (
        <div className="relative grid grid-cols-5 h-screen">
            <div className="col-span-2 flex items-center justify-center">
                
            </div>
            <div className="col-span-1"></div>
            <div className="col-span-2 flex items-center justify-center">
                <div className="border w-4/5 h-4/6 rounded-md flex flex-col justify-center">
                    <div className="max-w-md ">
                        <div>
                            <h1 className="text-2xl font-semibold">Welcome Back!</h1>
                            <p>Login to continue</p>
                        </div>
                        <div className="divide-y divide-gray-200">
                            <div className="py-8 text-base leading-6 space-y-4 text-gray-700 sm:text-lg sm:leading-7">
                                <div className="relative">
                                    <input autocomplete="off" id="email" name="email" type="text" className="peer placeholder-transparent h-10 w-full border-b-2 border-gray-300 text-gray-900 focus:outline-none focus:borer-rose-600" placeholder="Email address" />
                                    <label for="email" className="absolute left-0 -top-3.5 text-gray-600 text-sm peer-placeholder-shown:text-base peer-placeholder-shown:text-gray-440 peer-placeholder-shown:top-2 transition-all peer-focus:-top-3.5 peer-focus:text-gray-600 peer-focus:text-sm">Email Address</label>
                                </div>
                                <div className="relative">
                                    <input autocomplete="off" id="password" name="password" type="password" className="peer placeholder-transparent h-10 w-full border-b-2 border-gray-300 text-gray-900 focus:outline-none focus:borer-rose-600" placeholder="Password" />
                                    <label for="password" className="absolute left-0 -top-3.5 text-gray-600 text-sm peer-placeholder-shown:text-base peer-placeholder-shown:text-gray-440 peer-placeholder-shown:top-2 transition-all peer-focus:-top-3.5 peer-focus:text-gray-600 peer-focus:text-sm">Password</label>
                                </div>
                                <div className="relative">
                                    <button className="bg-blue-500 text-white rounded-md px-2 py-1">Submit</button>
                                </div>
                            </div>
                        </div>
                    </div>


                </div>
            </div>




            <div className="flex items-center justify-center absolute top-0 left-0 h-full w-3/5 bg-lime-200">
                <h1 className="gradient-text">GoSnap</h1>
            </div>
        </div>
    )
}

export default Auth