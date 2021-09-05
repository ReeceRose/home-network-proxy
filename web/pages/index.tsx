import { useEffect, useState } from 'react';

import Amplify, { Auth } from 'aws-amplify';
import awsconfig from '../aws-exports';
Amplify.configure(awsconfig);

import { withAuthenticator, AmplifySignOut } from '@aws-amplify/ui-react';
import axios from 'axios';
import { IPAddress } from '../interfaces/Index';

import Footer from '../components/Footer';
import ReportedIPAddress from '../components/Tables/ReportedIPAddress';
import CommandModal from '../components/Modal/Command';

const Home = (): JSX.Element => {
  const [addresses, setAddresses] = useState<IPAddress[]>();
  const [showModal, setShowModal] = useState(false);
  const [activeModalAddress, setActiveModalAddress] = useState<IPAddress>();

  useEffect(() => {
    async function fetchIPAddresses() {
      const user = await Auth.currentAuthenticatedUser();
      const token = user.signInUserSession.idToken.jwtToken;
      const response = await axios.get(
        'https://p7xsh0rld4.execute-api.us-east-1.amazonaws.com/production/ip/all',
        {
          headers: {
            Authorization: 'Bearer ' + token,
          },
        }
      );

      const addresses: IPAddress[] = JSON.parse(
        JSON.stringify(response.data.data)
      );
      setAddresses(addresses);
    }

    fetchIPAddresses();
  }, []);

  const deleteIP = (id: string) => {
    async function deleteIPAddress() {
      const user = await Auth.currentAuthenticatedUser();
      const token = user.signInUserSession.idToken.jwtToken;
      const response = await axios.delete(
        'https://p7xsh0rld4.execute-api.us-east-1.amazonaws.com/production/ip',
        {
          headers: {
            Authorization: 'Bearer ' + token,
          },
          data: JSON.stringify({ id: id }),
        }
      );
      if (response.status !== 200) {
        location.reload();
      }
      setAddresses(addresses?.filter((address) => address.Id.S !== id));
    }
    deleteIPAddress();
  };

  // const executeCommand = (id: string) => {
  //   async function lambda() {
  //     const user = await Auth.currentAuthenticatedUser();
  //     const token = user.signInUserSession.idToken.jwtToken;
  //     const response = await axios.delete(
  //       'https://p7xsh0rld4.execute-api.us-east-1.amazonaws.com/production/ssh',
  //       {
  //         headers: {
  //           Authorization: 'Bearer ' + token,
  //         },
  //         data: JSON.stringify({ id: id }),
  //       }
  //     );
  //     console.log(response);
  //   }
  //   lambda();
  // };

  const openCommandModal = (id: string) => {
    // TODO: open modal
    console.log('open modal for', id);
    setActiveModalAddress(addresses?.find((address) => address.Id.S == id));
    setShowModal(true);
  };

  return (
    <div className="relative flex flex-col w-full min-h-screen bg-gray-100">
      <div>
        <AmplifySignOut />
        <div className="flex-grow w-full px-4 m-24 mx-auto mb-0 md:px-10">
          <div className="flex flex-wrap">
            <div className="w-full px-4 mb-12 xl:mb-0">
              {addresses?.map((address) => (
                <ReportedIPAddress
                  key={address.Id.S}
                  address={address}
                  deleteIP={() => deleteIP(address.Id.S)}
                  openModal={() => openCommandModal(address.Id.S)}
                />
              ))}
            </div>
          </div>
          {showModal && activeModalAddress && (
            <CommandModal
              address={activeModalAddress}
              setModal={setShowModal}
            />
          )}
        </div>
        <Footer />
      </div>
    </div>
  );
};

export default withAuthenticator(Home);
