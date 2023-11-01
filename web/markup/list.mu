{{ define "list" }}
  {{ $list := index . "list" }}
  {{ range $i, $item := $list }}
  div p-3 bg-gray-900 flex space-x-3
    div
      input type=checkbox id=a{{$i}}
    div
      item {{$i}}
  {{ end }}
  {{ end }}
