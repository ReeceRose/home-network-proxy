import Link from 'next/link';

export default function Footer(): JSX.Element {
  return (
    <footer className="block py-4 absoulte">
      <div className="container px-24 mx-auto">
        <hr className="mb-4 border-gray-200 border-b-1" />
        <div className="flex flex-wrap items-center justify-center md:justify-between">
          <div className="w-full px-4 md:w-6/12">
            <div className="py-1 text-sm font-semibold text-center text-gray-500 md:text-left">
              Copyright © {new Date().getFullYear()}{' '}
              <Link href="https://reecerose.com/">
                <a
                  className="py-1 text-sm font-semibold text-gray-500 hover:text-gray-700"
                  target="_blank"
                  rel="noreferrer"
                >
                  Reece Rose
                </a>
              </Link>
            </div>
          </div>
          <div className="w-full px-4 md:w-6/12">
            <ul className="flex flex-wrap justify-center list-none md:justify-end">
              <li>
                <Link href="https://github.com/reecerose/home-network-proxy/">
                  <a
                    className="block px-3 py-1 text-sm font-semibold text-gray-600 hover:text-gray-800"
                    target="_blank"
                    rel="noreferrer"
                  >
                    GitHub
                  </a>
                </Link>
              </li>
            </ul>
          </div>
        </div>
      </div>
    </footer>
  );
}
