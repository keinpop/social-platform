import {
    LogOut,
    Settings,
    User,
  } from "lucide-react"
import { MiniUserIcon } from "@/components/ui/icons"
import LogoutButton from "@/app/dashboard/logout-button"
import Link from "next/link"
  
  import { Button } from "@/components/ui/button"
  import {
    DropdownMenu,
    DropdownMenuContent,
    DropdownMenuGroup,
    DropdownMenuItem,
    DropdownMenuLabel,
    DropdownMenuSeparator,
    DropdownMenuTrigger,
  } from "@/components/ui/dropdown-menu"
  
  export function DropdownMenuDemo() {
       
    return (
      <DropdownMenu>
        <DropdownMenuTrigger asChild>
          <div className="background-color: white">
            <button className="">
              <MiniUserIcon/>
            </button>
          </div>
        </DropdownMenuTrigger>
        <DropdownMenuContent className="w-56">
          <DropdownMenuLabel>Мой аккаунт</DropdownMenuLabel>
          <DropdownMenuSeparator />
          <DropdownMenuGroup>
            <DropdownMenuItem>
              <User className="mr-2 h-4 w-4" />
              <Link href="/dashboard/account">Домашняя страница</Link>
            </DropdownMenuItem>
            <DropdownMenuItem>
              <Settings className="mr-2 h-4 w-4" />
              <span>Настройки</span>
            </DropdownMenuItem>
          </DropdownMenuGroup>
          <DropdownMenuSeparator />
              <DropdownMenuItem>
                <LogOut  className="mr-2 h-4 w-4" />
                <LogoutButton/>
              </DropdownMenuItem>
        </DropdownMenuContent>
      </DropdownMenu>
      
    )
  }
  
