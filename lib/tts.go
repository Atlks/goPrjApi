package lib

//
//func TextToSpeech(text string) {
//	// 使用 gtranslate 库将文本转换为音频文件
//	audio, err := gtranslate.TranslateWithParams(
//		text,
//		gtranslate.TranslationParams{
//			From: "auto",
//			To:   "zh-CN",
//			//	Format: "mp3",
//		},
//	)
//	if err != nil {
//		log.Fatalf("Error translating text: %v", err)
//	}
//
//	// 保存音频文件
//	fileName := "output.mp3"
//	//err = ioutil.WriteFile(fileName, audio, 0644)
//	//if err != nil {
//	//	log.Fatalf("Error writing file: %v", err)
//	//}
//
//	fmt.Printf("Audio content written to file: %s\n", fileName)
//}
