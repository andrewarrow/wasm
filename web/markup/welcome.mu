div p-3 text-white flex
  div w-3/4 mx-auto
    div
      button id=b1 ml-9 border rounded bg-blue-600 text-white py-2 px-2
        this is button1
      button id=b2 ml-9 border rounded bg-blue-600 text-white py-2 px-2
        this is button2
    div bg-green-900 p-3 id=list
      {{ template "list" . }}
div opacity-0 bg-white transform translate-x-full transition-transform transition-opacity border-black border-2 fixed top-0 right-0 overflow-y-auto h-full id=form1
  {{ template "_form" . }}
