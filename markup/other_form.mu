{{ define "other_form" }}
  div p-3 bg-gray-900 flex space-x-3
    div
      other4 item 
    div
      input type=text placeholder=name
  div p-3 bg-gray-900 flex space-x-3
    div
      button id=cancel ml-9 border rounded bg-blue-600 text-white py-2 px-2
        Cancel
      button id=save ml-9 border rounded bg-blue-600 text-white py-2 px-2
        Save
  {{ end }}
