package main

import (
	"github.com/markbates/pkger"
	"github.com/markbates/pkger/pkging/mem"
)

var _ = pkger.Apply(mem.UnmarshalEmbed([]byte(`1f8b08000000000000ffec596b6fdbb8d2fe2b0b7e7d955017df22603f3469126477136c6ededd2e8a82a628893545b22495c42df2df5f90926c59bea6a73840718c201667e6e188d78723ce3740792a3488bf818c9abc9c1c635140995346a53e6268a2615e206cedefa90231808f9a280d114f66f0fdf9185e8a0ae081ab420a65fe442607f176671eb8410501312810e5c003ef050631001e78402a23a6f3964cc009e54dcd3b21ba80e5665c23837310ff0b8ec1470fdc1bc408888d2a492ddc11a4050731d056fa252192f084703c8b7f6935ba406a3a41866828a71951c00397e28232a2ad67dbece34c000f5863628b1f9bfe3bc056479b8df697f20c16a458c6ad0c21448cd877ed4029511af74e2632e001a16d3bdf13e91a3929536afb309919623d61514845b486294386b415d9572a9dcc0da29c28c8a836b582bcb8929a4923e605882a8f9580a9cc5d2b6a39691b138d1602c1cb6212f6fbc1c98a02526e88e28841923c2395e82e8c312a0dc50b4dbd366a695e5d219e9486b235265d4e0c230b4391f41782add79270af25b43ba073142c49617fb024f783b025775e69586b9c5efafec9b204e594be000f24c8a009d204ea2fac23c244d12737f4846391509eb58a10691eb4655b6bd05bd2508ed4acadc1faa92de6e4a52d7eb63bab254bb790e7e24b611b489412caf62c65c8b6272dccf23a465f89829aa329c14893651b66a24c5251f244cd8eb0288a92533383192a8d70b637c22b04e14f6fae273827d8543df9ce9ab0992b0d31a3841b2885369922d54cfee75e938999c9b78ee07a570a3dff103f1ae7a4406f7595102c922e7beeaee6763326d28837576502edfd427184d3952594648a7e46ec09c1cfcfe6c89d162d6baab930349dcd0bcbe64c4cca34454cc09ca8ce04e648e7140b25618ed91693fb47daec8248a474b7a76b50ca0de62e98c688f33d604661d11db155981153c2b7812ce1ecd10107dba3690eb7e6a59f0b419575f185bd6cb640455246b0d906b1fbb1c318f6b836084f21c1b9804fbded5658d02461e4197597c51c9889a2101c36c7fa160816acbb31ba902a70d802b047a1e8842b8c4ea0fcb24e07054dd6ea3556a8e8064719c58223aaa05442126568b7436f0aaee6477b81e4763f0ba8415322385c0bcd89443fc80d4c2933a47b94ac8f0ef78a20493121c9f7c59a1b70da2462a57dc670c77e761da109231bec5423633a0c57508373c258eea6c3a8129bb2bba425618c184a94756244c17e7054dc42696484a23013c765d95da29aaa526ae236832a3beeb54c8308a29428b1d1b03ac89511afb073a5fffc8cec727826c8e4441588af03c93a7c5a313c51d9eddd136233c490e3047bb010258560eb3129d2c69042d61f0199608867c74265f00536d13d2eb699a03b4589321b305ae7d01055508e5817c2898134e1a8abd7330d4beee2dd25bdfdfe809ae0521138a10955255b699ac31885b84e852ad65a4b4e6d80e15c6c05f0c6839c66c794c3192ad8f1930de2ed49523f205638723ae3d6ac7dc0d698ce79c1ce0762595b8565d916d3c268e14672aee2c41885f09223a15d54dd5649c1585bee56a90f2b46979ba429cf184919cdf2a5b7ea99c688b1661216d446b4a90e09f7314905a4a2fea4aad4053279fd80139a2d8a6ea3ba72fd395550b7acec031625335422d773a7f8520a43121785d44cc389a97e616e8c6c15dd4f33444bcaba61735dd38b5a67675a2ae13e64ad5c2a567d35db1f485e08ae4a65156ec8aa67f661099cd4723db0ae949117392f403de306d9d15325375567eb12c42e406ca4844ccaac25cf471b1951b8efda154b3dcc2b7a3db3cdafd7d022ecd246519e39d38ce3fab1705fcf36f040b579d0e459d18a342bc562312fc9552c66b555f7ea6db328c1d2a4c160591e3951a3d4e29e084f8425fc55dec039c2390a7d29d82c88fcfe0eb47bd81db62faef920de062ed513696e1bb6e0f269926e472c367cddab3de1f34b822de01d03649776c2b5fd2f88d628dbe46e69ef64a5dbb03b71528997d90e60087389f0740baa3e033699dd11bec16e8f8a8a45d759371c181ba1ed736323a8737cecc4b953e4a3071e8836f37b455e3256a9e6378995ea5a24b691f137b0dfe5ea35a2bcb9f2dc75697b29ae45b2196103a242240e38264a5377731a1c0711787d7df5405a357dfbe5710c0b9a2964a8e00e4b792aec33210651e654bcbe0f5ee03ca0e95702e29381070a4b217118f486bd512fe8054ef3c9514c0c423ff48ffcf028881ec220ee4571141d07416f3808fb7ef87f7e10fb766b51fd29b143518d8a25c5f81b784f9e401c0c86c36118fa6e9c899547c3a1076e18e55310471eb8e202c4e149d4eb477d0f3cd204c47d3ff0c0a52dd97a7789f5e37be04f947cc299f8e483f85fdf737f1f3df0ce36544be22ed1efed23e88f8293281cf4fb1eb8d156d33f89a228180ea3570f5cafc50fc2c120881afcbc83af1e385b870fa351300c070d3e0cfd6030e88d46af1e38a5cae46bdfd11b0d8351d8d4199c0c037f140683570fdccfe7e294093cd5aebba76c5acd51cfb7960b86b2ca7049b87bfea1a5fd4cb5c5dbba68c7e5a35b3a39513f6356616b13de9a7058c401fb67172a4e682ef11b0f3f3d99d4e1946dd734db7baadbccf25add7d371db5eb8d9b85c34525f7b67d090b4e8430c7fa0bdbce5c7354c35b41d06f88ab6749620763f97ed8eb0da3d16885b152c4f40eca8aa260349a5356505356e48f7ac361ef0752562f8aa261430fbd41d00ffbe1c9c99e9455f76f136175d0a3308c7a437f2b5d75fc47fdde70389a93959b8086ad46ff0db6824bc7d71271b5c8a98d59a5a805172d48a75acf35b1d4cb619959da0452a1d7d2c58214fe877962b1a1e784016e83df2e6ec77717571777a7b7d3978babcb0bfde1efdf7c3c7b27cee8bb0cfd759b3dfae3fbfbc78beb3fde89eceaf266f6e1af0bffc35f77e93f7fdf66e373767337be79bc7d0c6eaf2e92dfc6e759ad937f3e4c1fb387e9c9e3d5797ffc70fe9cdd9d8f2f6fc7e3ebf1d9693989b0f8fdfef9f7abb3d3e77ffe667a125e4cafcee5c383df3f3ba3d3e159f6ebaf600795558b93fc8c67e921437fc8d01f32f4870cfd21437fc8d01f32f4870cfd21437fc8d01f32f4870cfd21437fc8d01f32f4870cfd21437fc8d01f32f4870cfd2143ffbd19faff070000ffff010000ffffa91ec429df350000`)))