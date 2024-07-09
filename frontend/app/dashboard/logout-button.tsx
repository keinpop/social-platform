'use client';

import { logout } from '@/app/auth/01-auth';
import { LogOutIcon } from '@/components/ui/icons';
export default function LogoutButton() {
  return (
    <button
      className="flex items-center gap-3 rounded-lg px-0 py-0 text-sm font-weight-300 text-blue-950 transition-all hover:text-gray-900"
      onClick={async () => {
        await logout();
      }}
    >
      Logout
    </button>
  );
}
